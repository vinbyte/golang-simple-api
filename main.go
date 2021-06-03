package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vinbyte/golang-simple-api/db"
	"github.com/vinbyte/golang-simple-api/middleware"
	"github.com/vinbyte/golang-simple-api/router"
	"github.com/vinbyte/golang-simple-api/students"
)

func init() {
	_ = godotenv.Load()
}

func main() {
	//init db
	db := db.InitMysql()

	//setup framework
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	//init students package
	s := students.New(db)

	//setup middleware
	m := middleware.New()
	server.Use(m.CORS())

	//init router
	router.Init(server, s)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + os.Getenv("PORT"))
}
