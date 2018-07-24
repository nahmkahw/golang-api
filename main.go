package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/student/:id", getStudent)
	router.POST("/student", addStudent)
	router.PUT("/student", updateStudent)
	router.DELETE("/student", deleteStudent)

	router.Run(":9000")
}
