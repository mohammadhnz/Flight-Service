package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.Connect()
	routes.UserRoute(router)
	router.Run()
}
