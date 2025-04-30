package main

import (
	"crud.Restapi/crud/db"
	"crud.Restapi/crud/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.Registerroutes(server)
	server.Run(":8080")
}
