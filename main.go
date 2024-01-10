package main

import (
	"belajar-golang-db/controllers/itemcontroller"
	"belajar-golang-db/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	models.ConnectDB()

	r.Use(func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || !checkAuth(username, password) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		c.Next()
	})

	r.GET("/api/items", itemcontroller.Index)
	r.GET("/api/item/:id", itemcontroller.Show)
	r.POST("/api/item", itemcontroller.Create)
	r.PUT("/api/item/:id", itemcontroller.Update)
	r.DELETE("/api/item", itemcontroller.Delete)

	r.Run()
}

func checkAuth(username, password string) bool {
	return username == "root" && password == "password"
}
