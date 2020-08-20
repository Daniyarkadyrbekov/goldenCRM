package main

import (
	"html/template"
	"log"
	"math/rand"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"time"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gin-gonic/gin"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/goldenCRM.git/lib/handlers"
	"github.com/goldenCRM.git/lib/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
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

	database, err := getDatabase()
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

	//TODO: make main handler
	router.GET("/", func(c *gin.Context) {
		u := models.NewUser("Кадырбеков", "Данияр")
		flats := make([]models.Flat, 0)
		database.Find(&flats)

		c.HTML(200, "index.html", gin.H{
			"user":  &u,
			"flats": flats,
		})
	})

	router.GET("/flat/info", handlers.FlatInfo(l, database))
	router.POST("/flat/add", handlers.FlatNew(l, database))
	router.POST("/flat/search", handlers.FlatSearch(l, database))

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

func getDatabase() (database *gorm.DB, err error) {

	databaseUrl := os.Getenv("DATABASE_URL")
	if databaseUrl != "" {
		database, err = gorm.Open("postgres", databaseUrl)
		if err != nil {
			err = errors.Wrap(err, "creating postgres conn")
			return nil, err
		}

	} else {
		q := url.Values{}
		q.Set("sslmode", "disable")

		connUrl := &url.URL{
			Scheme:   "postgres",
			Host:     net.JoinHostPort("0.0.0.0", "5432"),
			User:     url.UserPassword("admin", "123"),
			Path:     url.QueryEscape("files"),
			RawQuery: q.Encode(),
		}

		database, err = gorm.Open("postgres", connUrl.String())
		if err != nil {
			err = errors.Wrap(err, "creating local postgres conn")
			return nil, err
		}
	}

	database.AutoMigrate(&models.Flat{}, &models.Owner{})
	database.Model(&models.Owner{}).AddForeignKey("id", "owners(id)", "CASCADE", "CASCADE")

	return
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
