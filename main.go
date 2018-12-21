package main

import (
	"github.com/501army/golang-simple-api/config"
	"github.com/501army/golang-simple-api/controllers"
	"github.com/501army/golang-simple-api/utils/db"
	"github.com/gin-gonic/gin"
)

var conf = config.ReadConfig()

func main() {
	welcomeController := controllers.NewWelcomeController()
	nameController := controllers.NewNameController()
	peopleController := new(controllers.PeopleController)

	db.Init()
	peopleController.Create()

	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", welcomeController.Welcome)
		v1.GET("/name", nameController.Name)
		v1.GET("/peoples", peopleController.FetchAll)
	}

	router.Run(":" + conf.Port)
}
