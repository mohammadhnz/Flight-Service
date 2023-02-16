package routes

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/all", controller.All)
	router.GET("/user_info", middleware.RequiredAuth, controller.UserInfo)
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)
	router.POST("/signout", controller.SignOut)
}
