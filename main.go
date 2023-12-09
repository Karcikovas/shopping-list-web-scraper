package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shoppingListWebSraper/scrapper"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.POST("/create-shopping-list", func(c *gin.Context) {
		var newTargetWebsites []scrapper.Website

		if err := c.BindJSON(&newTargetWebsites); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		data := scrapper.Scrap(newTargetWebsites)

		c.IndentedJSON(http.StatusOK, gin.H{"data": data})
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
