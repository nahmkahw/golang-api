package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	jwt_lib "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Username string
	Fullname string
	Token    string
	Role     string
}

var VerifyKey, SignKey []byte

func login(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user                  User
			Mysupersecretpassword = "nahmkahw@gmail.com"
		)

		username := c.PostForm("username")
		password := c.PostForm("password")

		fmt.Println(username, password)

		//token1 := uuid.NewV4()
		token := jwt_lib.New(jwt_lib.GetSigningMethod("HS256"))
		// Set some claims
		token.Claims = jwt_lib.MapClaims{
			"Id":  "nahmkahw",
			"exp": time.Now().Add(time.Hour * 1).Unix(),
		}
		// Sign and get the complete encoded token as a string
		tokenString, err := token.SignedString([]byte(Mysupersecretpassword))

		if err != nil {
			c.JSON(500, gin.H{"error": "Could not generate token"})
		}
		_, err = db.Exec("update user set token = ? where username = ? and password = ?", tokenString, username, password)
		if err != nil {
			c.JSON(500, gin.H{"error": "Could not update"})
		}
		sqlStatement := `select username,name,token,role from user where username = ?`
		row := db.QueryRow(sqlStatement, username)
		err = row.Scan(&user.Username, &user.Fullname, &user.Token, &user.Role)
		if err != nil {
			// If no results send null
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.Writer.Header().Set("Content-Type", "application/json")
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Max-Age", "86400")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

			c.JSON(http.StatusOK, gin.H{"token": tokenString})
		}
	}
}

func logout(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.PostForm("token")

		stmt, err := db.Prepare("UPDATE user set token = 'log-out' WHERE token = ?")
		defer stmt.Close()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		_, err = stmt.Exec(token)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
		} else {
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Set("Access-Control-Allow-Methods", "POST")
			c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			c.Next()
			c.JSON(http.StatusOK, gin.H{
				"message": fmt.Sprintf("successfully update token: %s", token),
				"error":   nil,
			})
		}
	}
}
