package main

import (
	"github.com/gin-contrib/cors"
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
	r.Use(cors.Default())

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.ValidateJwt, controllers.ValidateJwt)

	//Routes for Shows
	r.GET("/shows", middleware.ValidateJwt, controllers.GetShows)
	r.GET("/shows/:id", middleware.ValidateJwt, controllers.GetOneShow)
	r.POST("/shows", middleware.ValidateJwt, controllers.AddShows)
	r.POST("/watched", middleware.ValidateJwt, controllers.AddWatchedShows)
	r.GET("/watched/:id", middleware.ValidateJwt, controllers.GetWatched)

	//Routes for Foods
	r.GET("/foods", middleware.ValidateJwt, controllers.GetFoods)
	r.GET("/foods/:id", middleware.ValidateJwt, controllers.GetOneFood)
	r.POST("/foods", middleware.ValidateJwt, controllers.AddFood)

	r.Run() // listen and serve on env files
}
