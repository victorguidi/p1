package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/victorguidi/initializers"
	"github.com/victorguidi/models"
	"net/http"
)

func GetFoods(c *gin.Context) {

	// get all the Movies and Shows in the DB
	var foods []models.Foods
	result := initializers.DB.Find(&foods)

	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Couldn't get the data",
		})
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"foods": foods,
	})
}

func GetOneFood(c *gin.Context) {

	// get the id from the param
	id := c.Param("id")

	// get all the Movies and Shows in the DB
	var foods models.Foods
	result := initializers.DB.First(&foods, id)

	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Couldn't get the data",
		})
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"food": foods,
	})
}

func AddFood(c *gin.Context) {
	var body struct {
		Name    string
		Country string
		Type    string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// create the user
	food := models.Foods{Name: body.Name, Country: body.Country, Type: body.Type}

	result := initializers.DB.Create(&food)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create show",
		})

		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{})

}
