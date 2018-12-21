package controllers

import (
	"github.com/gin-gonic/gin"
)

// NameController is
type NameController struct{}

// NewNameController is
func NewNameController() *NameController {
	return &NameController{}
}

// Name is
func (w *NameController) Name(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"name":    "Gavinda Kinandana",
	})
}
