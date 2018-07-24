package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getStudent(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "get id %s", id)
}

func addStudent(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("Add id: %s; name: %s; message: %s", id, name, message)
	c.String(http.StatusOK, "Add id: %s; name: %s; message: %s", id, name, message)
}

func updateStudent(c *gin.Context) {
	id := c.Query("id")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("Update id: %s; name: %s; message: %s", id, name, message)
	c.String(http.StatusOK, "Update id: %s; name: %s; message: %s", id, name, message)
}

func deleteStudent(c *gin.Context) {
	id := c.Param("id")
	c.String(http.StatusOK, "Delete id %s", id)
}
