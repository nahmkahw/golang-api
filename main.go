package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()
	// Simple group: public
	routerpublic := router.Group("/public")
	{
		routerpublic.POST("/login", login)
		routerpublic.POST("/logout", logout)
	}

	// Simple group: private
	routerstudent := router.Group("/private")
	{
		routerstudent.GET("/student/:id", getStudent)
		routerstudent.POST("/student", addStudent)
		routerstudent.PUT("/student", updateStudent)
		routerstudent.DELETE("/student", deleteStudent)
	}

	router.Run(":9000")

}
