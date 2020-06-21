package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goldenCRM.git/lib/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func FlatInfo(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		ids, ok := params["ID"]
		if !ok || len(ids) != 1 {
			c.String(500, fmt.Sprintf("params = %v\n", params))
		}
		id, err := strconv.Atoi(ids[0])
		if err != nil {
			c.String(500, err.Error())
		}
		flat := models.Flat{}
		database.Where("flat_id = ?", id).First(&flat)
		c.JSON(200, flat)
	}
}

func FlatNew(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		flat, err := getFlatFromForm(c)
		if err != nil {
			l.Error("getting flat form testForm", zap.Error(err))
			c.String(500, "failed")
		}
		database.Create(&flat)
		c.Redirect(http.StatusFound, "/")
	}
}

func getFlatFromForm(c *gin.Context) (models.Flat, error) {

	area, ok := c.GetPostForm("InputArea")
	if !ok {
		return models.Flat{}, errors.New("no InputAddress in form")
	}
	landMark, ok := c.GetPostForm("InputLandMark")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	address, ok := c.GetPostForm("InputAddress")
	if !ok {
		return models.Flat{}, errors.New("no InputHomeNumber in form")
	}
	homeNumber, ok := c.GetPostForm("InputHomeNumber")
	if !ok {
		return models.Flat{}, errors.New("no InputHomeNumber in form")
	}
	flatNumber, ok := c.GetPostForm("InputFlatNumber")
	if !ok {
		return models.Flat{}, errors.New("no InputFlatNumber in form")
	}
	priceMin, ok := c.GetPostForm("InputPriceMin")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	priceMax, ok := c.GetPostForm("InputPriceMax")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	roomsCount, ok := c.GetPostForm("InputRoomsCount")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	roomsType, ok := c.GetPostForm("InputRoomsType")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	floor, ok := c.GetPostForm("InputFloor")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	floorsCount, ok := c.GetPostForm("InputFloorsCount")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	flatType, ok := c.GetPostForm("InputFlatType")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	square, ok := c.GetPostForm("InputSquare")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	state, ok := c.GetPostForm("InputState")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	toilet, ok := c.GetPostForm("InputToilet")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	toiletCount, ok := c.GetPostForm("InputToiletCount")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	buildYear, ok := c.GetPostForm("InputBuildYear")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	isCornerStr, ok := c.GetPostForm("inputIsCorner")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}
	description, ok := c.GetPostForm("InputDescription")
	if !ok {
		return models.Flat{}, errors.New("no InputArea in form")
	}

	flat, err := models.NewFlat(
		area,
		landMark,
		address,
		homeNumber,
		flatNumber,
		priceMin,
		priceMax,
		roomsCount,
		roomsType,
		floor,
		floorsCount,
		square,
		flatType,
		state,
		toilet,
		toiletCount,
		buildYear,
		isCornerStr == "on",
		description)

	if err != nil {
		return models.Flat{}, err
	}

	return flat, nil
}
