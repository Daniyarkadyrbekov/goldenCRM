package handlers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/goldenCRM.git/lib/models"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func FlatInfo(l *zap.Logger, database *gorm.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()
		ID, ok := params["ID"]
		if !ok || len(ID) != 1 {
			c.String(500, fmt.Sprintf("params = %v\n", params))
		}
		c.String(200, "Stub page with info of flat with ID = "+ID[0])
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
		//if err != nil {
		//	l.Error("adding flat to db err", zap.Error(err))
		//	c.String(500, "failed")
		//}
		c.Redirect(http.StatusFound, "/")
	}
}

func getFlatFromForm(c *gin.Context) (models.Flat, error) {

	/*name="InputArea"
	name="InputAddress"
	name="InputLandMark"
	name="InputHomeNumber"
	name="InputFlatNumber"
	name="InputPriceMin"
	name="InputPriceMax"
	name="InputRoomsCount"
	name="InputRoomsType"
	name="InputFloor"
	name="InputFloor"
	name="InputFlatType"
	name="InputSquare"
	name="InputToilet"
	name="InputToiletCount"
	name="InputBuildYear"
	name="inputIsCorner"
	name="InputDescription"*/

	//area, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}

	address, ok := c.GetPostForm("InputAddress")
	if !ok {
		return models.Flat{}, errors.New("no InputAddress in form")
	}
	//landMark, ok := c.GetPostForm("InputLandMark")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//homeNumber, ok := c.GetPostForm("InputHomeNumber")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputHomeNumber in form")
	//}
	//flatNumber, ok := c.GetPostForm("InputFlatNumber")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputFlatNumber in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}
	//inputArea, ok := c.GetPostForm("InputArea")
	//if !ok {
	//	return models.Flat{}, errors.New("no InputArea in form")
	//}

	flat, err := models.NewFlat(
		"area",
		"landMark",
		address,
		"1",
		"1",
		"1",
		"1",
		"1",
		"",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		"1",
		false,
		"1",
		[]string{},
		map[string]string{})

	if err != nil {
		return models.Flat{}, err
	}

	return flat, nil
}
