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
	l = l.With(zap.String("method", "info"))
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		ids, ok := params["ID"]
		if !ok || len(ids) != 1 {
			l.Error("id not defined")
			c.String(500, fmt.Sprintf("params = %v\n", params))
			return
		}
		id, err := strconv.Atoi(ids[0])
		if err != nil {
			l.Error("converting id err", zap.Error(err))
			c.String(500, err.Error())
			return
		}
		flat := models.Flat{}
		database.Where("flat_id = ?", id).First(&flat)

		owners := make([]models.Owner, 0)
		database.Model(&flat).Association("Owners").Find(&owners)
		flat.Owners = owners

		c.HTML(200, "flat.html", gin.H{
			"flat": flat,
		})
	}
}

func FlatAdd(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {

	l = l.With(zap.String("method", "FlatAdd"))
	return func(c *gin.Context) {
		flat, err := getFlatFromForm(c, true)
		if err != nil {
			l.Error("getting flat form testForm", zap.Error(err))
			c.String(500, "failed")
			return
		}

		database.Create(&flat)
		c.Redirect(http.StatusFound, "/")
	}
}

func FlatSearch(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		flat, err := getFlatFromForm(c, false)
		if err != nil {
			l.Error("getting flat form testForm", zap.Error(err))
			c.String(500, "failed")
			return
		}

		u := models.NewUser("Кадырбеков", "Данияр")
		flats := make([]models.Flat, 0)

		owners := flat.Owners
		flat.Owners = nil
		complexCondition := getComplexCondition(&flat)
		database.Where(&flat).Where(complexCondition).Preload("Owners").Find(&flats)

		//filter flats by ownerPhoneNumber
		if len(owners) == 1 && owners[0].Phone != "" {
			result := make([]models.Flat, 0, len(flats))
			for _, f := range flats {
				matched := false
				for _, owner := range f.Owners {
					if owners[0].Phone == owner.Phone {
						matched = true
						break
					}
				}
				if matched {
					result = append(result, f)
				}
			}
			flats = result
		}

		c.HTML(200, "index.html", gin.H{
			"user":  &u,
			"flats": flats,
		})
	}
}

func getFlatFromForm(c *gin.Context, requiredFieldsCheck bool) (models.Flat, error) {

	landMark, ok := c.GetPostForm("InputLandMark")
	if requiredFieldsCheck && (!ok || landMark == "") {
		return models.Flat{}, errors.New("no InputLandMark in form")
	}

	// optional fields
	area, _ := c.GetPostForm("InputArea")
	address, _ := c.GetPostForm("InputAddress")
	homeNumber, _ := c.GetPostForm("InputHomeNumber")
	flatNumber, _ := c.GetPostForm("InputFlatNumber")
	priceMin, _ := c.GetPostForm("InputPriceMin")
	priceMax, _ := c.GetPostForm("InputPriceMax")
	roomsCount, _ := c.GetPostForm("InputRoomsCount")
	roomsType, _ := c.GetPostForm("InputRoomsType")
	floor, _ := c.GetPostForm("InputFloor")
	floorsCount, _ := c.GetPostForm("InputFloorsCount")
	flatType, _ := c.GetPostForm("InputFlatType")
	square, _ := c.GetPostForm("InputSquare")
	state, _ := c.GetPostForm("InputState")
	toilet, _ := c.GetPostForm("InputToilet")
	toiletCount, _ := c.GetPostForm("InputToiletCount")
	buildYear, _ := c.GetPostForm("InputBuildYear")
	isCornerStr, _ := c.GetPostForm("inputIsCorner")
	isSeparatedStr, _ := c.GetPostForm("inputIsSeparated")
	description, _ := c.GetPostForm("InputDescription")

	owners := getOwnersFromForm(c)

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
		isSeparatedStr == "on",
		description,
		owners)

	if err != nil {
		return models.Flat{}, err
	}

	return flat, nil
}

func getOwnersFromForm(c *gin.Context) []models.Owner {
	owners := make([]models.Owner, 0)

	counter := 0
	for {
		cStr := strconv.Itoa(counter)
		counter++

		ownerName, nameExists := c.GetPostForm("InputOwnerName" + cStr)
		ownerPhone, phoneExists := c.GetPostForm("InputOwnerPhoneNumber" + cStr)
		if !nameExists && !phoneExists {
			return owners
		}

		owners = append(owners, models.Owner{
			Name:  ownerName,
			Phone: ownerPhone,
		})
	}
}
