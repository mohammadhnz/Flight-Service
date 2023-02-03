package routes

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/user_info", controller.UserInfo)
	router.POST("/sign-up", controller.SignUp)
	router.POST("/sign-in", controller.SignIn)
	router.POST("/sign-out", controller.SignOut)
}
