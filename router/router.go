package router

import (
	"github.com/gin-gonic/gin"
	"github.com/vinbyte/golang-simple-api/students"
)

func Init(r *gin.Engine, s students.Student) {
	v1 := r.Group("v1")
	{
		v1.GET("/students", s.StudentList)
		v1.POST("/student", s.StudentAdd)
	}
}
