package routes

import (
	"crud.Restapi/crud/middlewares"
	"github.com/gin-gonic/gin"
)

func Registerroutes(server *gin.Engine) {
	server.GET("/events", Events)
	server.GET("/events/:id", Getevent)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", CreateEvents)
	authenticated.PUT("/events/:id", UpdateEvent)
	authenticated.DELETE("/events/:id", DeleteEvent)
	server.POST("/signup", createuser)
	server.POST("/login", login)
}
