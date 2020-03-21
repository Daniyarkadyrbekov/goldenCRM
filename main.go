package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"html/template"

	"github.com/gin-gonic/gin"
)

var tpl = template.Must(template.ParseFiles([]string{"pages/index.html", "pages/base_header.html"}...))

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.LoadHTMLGlob("pages/*.html")
	router.Use(gin.Logger())

	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/sources/:ext/:fileName", func(c *gin.Context) {
		ext := c.Param("ext")
		fileName := c.Param("fileName")
		file := fmt.Sprintf("pages/sources/%s/%s", ext, fileName)
		c.File(file)
	})

	router.Run(":" + port)
}
