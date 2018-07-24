package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Disable Console Color
	// gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/student", getStudent)
	router.POST("/student", addStudent)
	router.PUT("/student", updateStudent)
	router.DELETE("/student", deleteStudent)

	router.Run(":9000")
}

func getStudent(c *gin.Context) {
	c.String(http.StatusOK, "get Student")
}

func addStudent(c *gin.Context) {
	c.String(http.StatusOK, "add Student")
}

func updateStudent(c *gin.Context) {
	c.String(http.StatusOK, "update Student")
}

func deleteStudent(c *gin.Context) {
	c.String(http.StatusOK, "delete Student")
}
