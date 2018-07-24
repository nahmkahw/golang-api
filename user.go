package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func logout(c *gin.Context) {
	c.String(http.StatusOK, "logout")
}
