package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	courseAPI := r.Group("/api/courses")
	{
		courseAPI.POST("/", createCourse)
		courseAPI.GET("/", getCourses)
		courseAPI.GET("/:id", getCourseByID)
		courseAPI.PUT("/:id", updateCourse)
		courseAPI.DELETE("/:id", deleteCourse)
	}
	courseStudentsAPI := r.Group("/api/courses/:id/students")
	{
		courseStudentsAPI.GET("/", getCourseStudents)
		courseStudentsAPI.POST("/", addStudentToCourse)
		courseStudentsAPI.DELETE("/:studentID", removeStudentFromCourse)
	}
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}
