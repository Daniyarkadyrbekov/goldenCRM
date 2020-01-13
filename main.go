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
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		//c.HTML(http.StatusOK, "index.tmpl.html", nil)
		c.String(200, "hello world\n")
	})

	router.Run(":" + port)
}
