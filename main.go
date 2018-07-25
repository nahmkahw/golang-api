package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	_ "github.com/go-sql-driver/mysql"
)

type Tag struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func main() {

	// db.SetMaxIdleConns(3)
	// db.SetMaxOpenConns(3)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		db, err := sql.Open("mysql", "gouser:golang@tcp(202.41.160.101:4406)/barcode?charset=utf8")
		checkErr(err)
		defer db.Close()
		var (
			tag  Tag
			tags []Tag
		)
		results, err := db.Query("SELECT username, name FROM user")
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		for results.Next() {
			// for each row, scan the result into our tag composite object
			err = results.Scan(&tag.ID, &tag.Name)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}
			// and then print out the tag's Name attribute
			//log.Printf(tag.Name)
			tags = append(tags, tag)
		}
		c.JSON(200, gin.H{
			"message": tags,
		})
	})
	r.Run(":9000") // listen and serve on 0.0.0.0:8080

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
