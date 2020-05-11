package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/goldenCRM.git/lib/storage"
	"github.com/goldenCRM.git/lib/storage/postgres"
	"github.com/spf13/viper"

	"github.com/goldenCRM.git/lib/storage/mock"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/goldenCRM.git/lib/models"
	"go.uber.org/zap"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {

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

	var database storage.Storage
	if false {
		database = mock.New()
	} else {
		v := viper.New()
		v.SetDefault("host", "0.0.0.0")
		v.SetDefault("port", "5432")
		v.SetDefault("name", "files")
		v.SetDefault("user", "admin")
		v.SetDefault("password", "123")
		v.SetDefault("ssl-mode", "disable")
		v.SetDefault("schema", "schema")
		v.SetDefault("health-check", time.Second)
		v.SetDefault("max-connections", 10)
		conf, err := postgres.NewConfig(v)
		if err != nil {
			l.Fatal("get config for postgres", zap.Error(err))
		}
		database, err = postgres.New(context.Background(), conf)
		if err != nil {
			l.Fatal("creating postgres", zap.Error(err))
		}
	}
	router := gin.New()
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
		u := models.NewUser("Кадырбеков", "Данияр")
		flats, err := database.List()
		if err != nil {
			l.Error("getting list err", zap.Error(err))
			c.String(500, "getting list err")
		}

		c.HTML(200, "index.html", gin.H{
			"user":  &u,
			"flats": flats,
		})
	})

	//TODO: route static files with static router method
	router.GET("/sources/:ext/:fileName", func(c *gin.Context) {
		ext := c.Param("ext")
		fileName := c.Param("fileName")
		file := fmt.Sprintf("pages/sources/%s/%s", ext, fileName)
		c.File(file)
	})

	router.POST("/flat/new", func(c *gin.Context) {
		flat, err := getFlatFromTestForm(c)
		if err != nil {
			l.Error("getting flat form testForm", zap.Error(err))
			c.String(500, "failed")
		}
		err = database.Add(flat)
		if err != nil {
			l.Error("adding flat to db err", zap.Error(err))
			c.String(500, "failed")
		}
		c.String(200, "success")
	})

	err = router.Run(":" + port)
	if err != nil {
		l.Error("closing server", zap.Error(err))
	}
}

func getFlatFromTestForm(c *gin.Context) (models.Flat, error) {
	//models.Flat{
	//	ID:          0,
	//	Street:      "",
	//	Home:        "",
	//	Structure:   0,
	//	FlatNumber:  0,
	//	State:       "",
	//	Floor:       0,
	//	IsCorner:    false,
	//	FlatType:    "",
	//	Description: "",
	//	PictureURLs: nil,
	//	Owner:       "",
	//}
	flat := models.NewFlat(c.PostForm("inputStreet"),
		"", 1, 1,
		models.Euro, 1, false,
		"", "", []string{""}, "")

	return flat, nil
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
