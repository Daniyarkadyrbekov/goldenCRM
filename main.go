package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

//var tpl = template.Must(template.ParseFiles([]string{"pages/index.html", "pages/base_header.html"}...))
//
//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	tpl.Execute(w, nil)
//}

func main() {

	t, err := template.ParseFiles("pages/templates/base_header.html", "pages/templates/index.html")
	if err != nil {
		panic(err)
	}

	//err = t.ExecuteTemplate(os.Stdout, "index.html", nil)
	//if err != nil {
	//	panic(err)
	//}

	cfg := zap.NewDevelopmentConfig()
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	l, err := cfg.Build()
	if err != nil {
		log.Fatal("error creating log", err)
	}
	l.Info("starting service")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.LoadHTMLGlob("pages/templates/*")
	router.Use(gin.Logger())

	tmplBox, err := rice.FindBox("pages/templates")
	if err != nil {
		l.Fatal("can't find templates box", zap.Error(err))
	}
	if tmplBox == nil {
		l.Fatal("tmplBox is nil")
	}

	err = initResources(l, router, tmplBox)
	if err != nil {
		l.Fatal("can't init resources", zap.Error(err))
	}

	router.GET("/", func(c *gin.Context) {
		//var b bytes.Buffer
		//foo := bufio.NewWriter(&b)
		//err := tpl.ExecuteTemplate(foo, "base_header", "")
		//if err != nil {
		//	c.Abort()
		//}
		var b bytes.Buffer
		err := t.ExecuteTemplate(&b, "base_header.html", nil)
		if err != nil {
			c.Abort()
		}
		//c.String(200, b.String())

		c.HTML(200, "index.html", gin.H{
			"base_header": b.String(),
		})
	})

	router.GET("/sources/:ext/:fileName", func(c *gin.Context) {
		ext := c.Param("ext")
		fileName := c.Param("fileName")
		file := fmt.Sprintf("pages/sources/%s/%s", ext, fileName)
		c.File(file)
	})

	err = router.Run(":" + port)
	if err != nil {
		l.Error("closing server", zap.Error(err))
	}
}

func initResources(l *zap.Logger, r *gin.Engine, templates *rice.Box) error {
	l.Debug("init resources")
	if templates != nil {
		l.Info("Serving embedded templates")
		var tmplt *template.Template
		err := templates.Walk("/", func(path string, info os.FileInfo, err error) error {
			path = filepath.Base(path)
			if err != nil {
				l.Error("WERR", zap.Error(err))
				return err
			}
			if info.IsDir() {
				return nil
			}

			templateString, err := templates.String(path)
			if err != nil {
				return err
			}
			if len(templateString) < 1 {
				return nil
			}

			l.Debug("template New", zap.String("path", path))
			if tmplt == nil {
				tmplt, err = template.New(path).Parse(templateString)
			} else {
				tmplt, err = tmplt.New(path).Parse(templateString)
			}

			return nil
		})
		if err != nil {
			return err
		}
		if tmplt == nil {
			l.Fatal("Templates are nil")
		}
		r.SetHTMLTemplate(tmplt)
	} else {
		l.Info("Serving local templates")
		r.LoadHTMLGlob("templates/*")
	}

	//if statics != nil {
	//	l.Info("Serving embedded statics")
	//	r.StaticFS("/static", statics.HTTPBox())
	//} else {
	//	l.Info("Serving local statics")
	//	r.Static("/static", "static")
	//}

	return nil
}
