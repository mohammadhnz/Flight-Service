package main

import (
	"awesomeProject/config"
	"awesomeProject/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.InitializeEnvVars()
	config.Connect()
}
func main() {
	router := gin.Default()
	routes.UserRoute(router)
	router.Run()
}
