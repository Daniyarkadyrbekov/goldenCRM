package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goldenCRM.git/lib/models"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func AdminMain(l *zap.Logger) func(c *gin.Context) {
	l = l.With(zap.String("method", "adminMain"))
	return func(c *gin.Context) {
		c.HTML(200, "admin_main.html", gin.H{})
	}
}

func AdminGetAddresses(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "getAddresses"))
	return func(c *gin.Context) {
		var addresses []models.Address
		database.Find(&addresses)

		c.HTML(200, "admin_addresses_list.html", gin.H{
			"addresses": addresses,
		})
	}
}

func AdminAddAddress(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "addAddress"))
	return func(c *gin.Context) {
		address, ok := c.GetPostForm("Address")
		if !ok {
			c.JSON(500, "не введен адресс")
		}
		addressModel := &models.Address{
			Address: address,
		}
		if err := database.Create(addressModel).Error; err != nil {
			l.Error("create address", zap.String("address", address), zap.Error(err))
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Redirect(http.StatusFound, "/auth/admin/addresses")
	}
}

func AdminDeleteAddress(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "deleteAddress"))
	return func(c *gin.Context) {
		id, ok := c.GetPostForm("ID")
		if !ok {
			c.JSON(500, "не введен адресс")
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			l.Error("id convert err", zap.String("id", id), zap.Error(err))
			c.String(http.StatusInternalServerError, err.Error())
		}
		idUint := uint(idInt)
		addressModel := &models.Address{
			Model: gorm.Model{ID: idUint},
		}
		if err := database.Delete(addressModel).Error; err != nil {
			l.Error("delete address", zap.Uint("id", idUint), zap.Error(err))
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Redirect(http.StatusFound, "/auth/admin/addresses")
	}
}

func AdminGetLandmarks(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "getLandmarks"))
	return func(c *gin.Context) {
		var landmarks []models.Landmark
		database.Find(&landmarks)

		c.HTML(200, "admin_landmarks_list.html", gin.H{
			"landmarks": landmarks,
		})
	}
}

func AdminAddLandmark(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "addLandmarks"))
	return func(c *gin.Context) {
		landmark, ok := c.GetPostForm("Landmark")
		if !ok {
			c.JSON(500, "не введен адресс")
		}
		landmarkModel := &models.Landmark{
			Landmark: landmark,
		}
		if err := database.Create(landmarkModel).Error; err != nil {
			l.Error("create landmarks", zap.String("landmark", landmark), zap.Error(err))
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Redirect(http.StatusFound, "/auth/admin/landmarks")
	}
}

func AdminDeleteLandmark(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	l = l.With(zap.String("method", "deleteLandmarks"))
	return func(c *gin.Context) {
		id, ok := c.GetPostForm("ID")
		if !ok {
			c.JSON(500, "не введен адресс")
		}
		idInt, err := strconv.Atoi(id)
		if err != nil {
			l.Error("id convert err", zap.String("id", id), zap.Error(err))
			c.String(http.StatusInternalServerError, err.Error())
		}
		idUint := uint(idInt)
		landmarkModel := &models.Landmark{
			Model: gorm.Model{ID: idUint},
		}
		if err := database.Delete(landmarkModel).Error; err != nil {
			l.Error("delete landmark", zap.Uint("id", idUint), zap.Error(err))
			c.String(http.StatusInternalServerError, err.Error())
		}

		c.Redirect(http.StatusFound, "/auth/admin/landmarks")
	}
}
