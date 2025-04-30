package routes

import (
	"net/http"

	models "crud.Restapi/crud/Models"
	"crud.Restapi/crud/utils"
	"github.com/gin-gonic/gin"
)

func createuser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Enter All fileds "})
		return
	}
	user.Id = 1

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Some Err occur in create user",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User Created",
	})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message": "Please fill Email and Password",
		})
		return
	}
	err = user.ValidateCrendentials()
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"Message": "Password might be wrong",
		})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.Id)
	if err != nil {
		context.JSON(http.StatusForbidden, gin.H{
			"Message": "Couldn't Generate JWT Token",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"Message": "Login Successfully", "Token": token,
	})
}
