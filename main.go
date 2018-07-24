package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

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
