package routes

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/all", controller.All)
	router.GET("/user_info", controller.UserInfo)
	router.POST("/signup", controller.SignUp)
	router.POST("/signin", controller.SignIn)
	router.POST("/signout", controller.SignOut)
}
