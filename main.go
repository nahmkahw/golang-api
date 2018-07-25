package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var Mysupersecretpassword = "nahmkahw@gmail.com"

func main() {

	db, err := sql.Open("mysql", "gouser:golang@tcp(202.41.160.101:4406)/barcode?charset=utf8")
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connect database success.")
	}
	defer db.Close()
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

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
		routerpublic.POST("/login", login(db))
		routerpublic.POST("/logout", logout(db))
	}

	// Simple group: private
	routerstudent := router.Group("/student")
	{
		routerstudent.GET("/", FetchAllStudent(db))
		routerstudent.GET("/:id", Fetchstudent(db))
		routerstudent.POST("/", Createstudent(db))
		routerstudent.PUT("/:id", Updatestudent(db))
		routerstudent.DELETE("/:id", Deletestudent(db))
	}

	router.Run(":9000")

}
