package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/goldenCRM.git/lib/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func MainPage(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "MainPage"))

	return func(c *gin.Context) {
		u := models.NewUser("Кадырбеков", "Данияр")
		flats := make([]models.Flat, 0)
		database.Find(&flats)

		addresses := make([]models.Address, 0)
		database.Find(&addresses)

		landmarks := make([]models.Landmark, 0)
		database.Find(&landmarks)

		c.HTML(200, "index.html", gin.H{
			"user":      &u,
			"flats":     flats,
			"addresses": addresses,
			"landmarks": landmarks,
		})
	}
}
