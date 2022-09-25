package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/victorguidi/initializers"
	"github.com/victorguidi/models"
)

func GetShows(c *gin.Context) {
	// get all the Movies and Shows in the DB
	var shows []models.Shows
	result := initializers.DB.Find(&shows)

	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Couldn't get the data",
		})
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"shows": shows,
	})
}

func GetOneShow(c *gin.Context) {

	// get the id from the param
	id := c.Param("id")

	// get all the Movies and Shows in the DB
	var shows models.Shows
	result := initializers.DB.First(&shows, id)

	if result.Error != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": "Couldn't get the data",
		})
	}

	// Return data
	c.JSON(http.StatusOK, gin.H{
		"show": shows,
	})

}

func AddShows(c *gin.Context) {
	var body struct {
		Name  string
		Type  string
		Class string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	// create the user
	show := models.Shows{Name: body.Name, Type: body.Type, Class: body.Class}

	result := initializers.DB.Create(&show)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create show",
		})

		return
	}

	// respond
	c.JSON(http.StatusOK, gin.H{})

}
