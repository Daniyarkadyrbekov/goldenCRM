package handlers

import (
	"fmt"
	"net/http"

	"github.com/goldenCRM.git/lib/storage"

	"github.com/goldenCRM.git/lib/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func FlatInfo(l *zap.Logger, database storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		ID, ok := params["ID"]
		if !ok || len(ID) != 1 {
			c.String(500, fmt.Sprintf("params = %v\n", params))
		}
		c.String(200, "Stub page with info of flat with ID = "+ID[0])
	}
}

func FlatNew(l *zap.Logger, database storage.Storage) func(c *gin.Context) {
	return func(c *gin.Context) {
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
		c.Redirect(http.StatusFound, "/")
	}
}

func getFlatFromTestForm(c *gin.Context) (models.Flat, error) {
	flat := models.NewFlat(c.PostForm("inputStreet"),
		"",
		1,
		1,
		models.Euro,
		1,
		false,
		"",
		"",
		[]string{""},
		"")

	return flat, nil
}
