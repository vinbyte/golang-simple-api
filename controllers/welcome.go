package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// WelcomeController is
type WelcomeController struct{}

// Welcome is
func (w *WelcomeController) Welcome(c *gin.Context) {
	c.String(http.StatusOK, "Hello there, welcome")
}
