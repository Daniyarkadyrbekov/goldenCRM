package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Printf("hello world!\n")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl.html", nil)
	})

	router.Run(":" + port)
}
