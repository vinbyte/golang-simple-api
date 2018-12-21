package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeController is
type WelcomeController struct{}

// NewWelcomeController is
func NewWelcomeController() *WelcomeController {
	return &WelcomeController{}
}

// Welcome is
func (w *WelcomeController) Welcome(c *gin.Context) {
	c.String(http.StatusOK, "Hello there, welcome")
}
