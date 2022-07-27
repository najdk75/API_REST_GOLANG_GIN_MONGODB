package routes

import (
	"github.com/gin-gonic/gin"
	"goAPI/controllers"
)

func Routes(router *gin.Engine) {

	router.POST("/create", controllers.CreateUser)

	router.GET("/user/:id", controllers.GetById)
	router.GET("/users/list", controllers.GetUsers)

	router.DELETE("/delete/user/:id", controllers.DeleteUser)
	router.DELETE("/delete/users", controllers.DeleteAll)

	router.POST("/login", controllers.LoginUser)
	router.PUT("/update/user/:id", controllers.UpdateUser)
}
