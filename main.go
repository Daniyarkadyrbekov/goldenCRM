package main

import (
	"context"
	"html/template"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/goldenCRM.git/lib/handlers"

	"github.com/pkg/errors"

	"github.com/goldenCRM.git/lib/storage/postgres"
	"github.com/spf13/viper"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	"github.com/goldenCRM.git/lib/models"
	"github.com/goldenCRM.git/lib/storage"
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

	database, err := getDatabase(l)
	if err != nil {
		l.Fatal(err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	err = initResources(l, router)
	if err != nil {
		l.Fatal("can't init resources", zap.Error(err))
	}

	//TODO: make main handler |
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

	router.GET("/flat/info", handlers.FlatInfo(l, database))
	router.POST("/flat/new", handlers.FlatNew(l, database))

	err = router.Run(":" + port)
	if err != nil {
		l.Error("closing server", zap.Error(err))
	}
}

func getBoxes() (tmplBox, sourcesBox *rice.Box, err error) {

	tmplBox, err = rice.FindBox("pages/templates")
	if err != nil {
		err = errors.Wrap(err, "can't find templates box")
		return
	}

	sourcesBox, err = rice.FindBox("pages/sources")
	if err != nil {
		err = errors.Wrap(err, "can't find sources box")
		return
	}

	if tmplBox == nil || sourcesBox == nil {
		err = errors.New("tmplBox is nil")
		return
	}

	return
}

func getDatabase(l *zap.Logger) (storage.Storage, error) {

	var database storage.Storage

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl != "" {
		db, err := postgres.New(context.Background(), databaseUrl)
		if err != nil {
			err = errors.Wrap(err, "creating postgres conn")
			return nil, err
		}
		database = db
	} else {
		v := viper.New()
		v.SetDefault("host", "0.0.0.0")
		v.SetDefault("port", "5432")
		v.SetDefault("name", "files")
		v.SetDefault("user", "admin")
		v.SetDefault("password", "123")
		v.SetDefault("ssl-mode", "disable")
		v.SetDefault("schema", "")
		v.SetDefault("health-check", time.Second)
		v.SetDefault("max-connections", 10)
		conf, err := postgres.NewConfig(v)
		if err != nil {
			err = errors.Wrap(err, "get config for postgres")
			return nil, err
		}
		database, err = postgres.New(context.Background(), conf.ConnURL())
		if err != nil {
			err = errors.Wrap(err, "creating local postgres conn")
			return nil, err
		}
	}

	return database, nil
}

func initResources(l *zap.Logger, r *gin.Engine) error {
	templates, sources, err := getBoxes()
	if err != nil {
		return errors.Wrap(err, "getting boxes err")
	}

	var tmplt *template.Template
	err = templates.Walk("/", func(path string, info os.FileInfo, err error) error {
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

	r.StaticFS("/sources", sources.HTTPBox())

	return nil
}
