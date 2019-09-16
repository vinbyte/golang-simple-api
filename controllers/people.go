package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/rabbitmeow/golang-simple-api/models"
)

// PeopleController is
type PeopleController struct{}

var peopleModel = new(models.PeopleModel)

// Create is
func (w *PeopleController) Create() {
	peopleModel.Init()
}

// FetchAll is
func (w *PeopleController) FetchAll(c *gin.Context) {
	list := peopleModel.GetAll()
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    list,
	})
}
