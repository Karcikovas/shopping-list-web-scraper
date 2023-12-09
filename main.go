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

	r.GET("/create-shopping-list", func(c *gin.Context) {
		user := c.Params.ByName("name")
		data := scrapper.Scrap()

		if len(data) != 0 {
			c.IndentedJSON(http.StatusOK, gin.H{"data": data})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
