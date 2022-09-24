package main

import (
	"github.com/gin-gonic/gin"
	"github.com/victorguidi/controllers"
	"github.com/victorguidi/initializers"
	"github.com/victorguidi/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDb()
}

func main() {
	r := gin.Default()
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.ValidateJwt, controllers.ValidateJwt)

	r.Run() // listen and serve on env files
}
