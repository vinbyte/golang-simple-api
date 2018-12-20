package main

import (
	"net/http"

	"github.com/501army/golang-simple-api/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", welcome)
		v1.GET("/name", name)
	}
	conf := config.ReadConfig()
	router.Run(":" + conf.Port)
}

func welcome(c *gin.Context) {
	c.String(http.StatusOK, "Hello there, welcome")
}

func name(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"name":    "Gavinda Kinandana",
	})
}
