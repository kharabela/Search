package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var items = []string{"Apple", "Banana", "Cherry", "Date", "Fig", "Grape", "Kiwi"}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	r.GET("/search", func(c *gin.Context) {
		query := strings.ToLower(c.Query("q"))
		var results []string
		for _, item := range items {
			if strings.Contains(strings.ToLower(item), query) {
				results = append(results, item)
			}
		}
		c.HTML(http.StatusOK, "results.html", gin.H{"Results": results})
	})

	r.Run(":8090")
}
