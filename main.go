package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Use(gin.Logger())

	router.GET("/main", func(c *gin.Context) {
		fmt.Printf("===>corePath\n")
		c.HTML(200, "./templates/index.tmpl.html", nil)
	})

	router.GET("/get/", func(c *gin.Context) {
		fmt.Printf("===>cssPath\n")
		//ext := c.Param("ext")
		//fileName := c.Param("fileName")
		//file := fmt.Sprintf("./sources/%s/%s", ext, fileName)
		//fmt.Printf("get Ext fileName = %s\n", file)
		c.HTML(200, "./templates/sources/css/bootstrap.min.css", nil)
	})

	router.Run(":" + port)
}
