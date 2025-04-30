package models

import (
	"errors"

	"crud.Restapi/crud/db"
	"crud.Restapi/crud/utils"
)

type User struct {
	Id       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user User) Save() error {

	query := `INSERT INTO users(email,password) VALUES (?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashpass, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashpass)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()
	if err != nil {
		return err
	}

	return nil
}

func (user *User) ValidateCrendentials() error {
	query := `SELECT id,password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)
	var retrivedpass string
	err := row.Scan(&user.Id, &retrivedpass)
	if err != nil {
		return err
	}
	passwordisvalid := utils.Checkpassword(user.Password, retrivedpass)
	if !passwordisvalid {
		return errors.New("invalid credentials")
	}
	return nil
}
