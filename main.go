package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/service"
)

func main() {
	srv := service.NewUserService()

	r := gin.Default()
	r.POST("/users", func(c *gin.Context) {

		var user service.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, _ = srv.Save(user)
		c.JSON(http.StatusCreated, user)
	})
	r.GET("/users", func(c *gin.Context) {
		users := srv.Find()
		c.JSON(http.StatusOK, users)
	})
	r.Run()
}
