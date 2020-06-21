package models

import (
	"database/sql"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/pkg/errors"
)

//Телефон (+7 по базе должно высвечиваться автоматически)
//Район
//Адрес
//дом №
//квартира №
//Цена объекта  от ………………….до …………..
//Кол -во комнат
//тип комнат
//Этаж / этажность
//Серия
//площадь М2
//Сан узел (Раздельный ,Во дворе,  Нет ,   Совмещенный)
//Кол -во сан узлов
//Год постройки
//Угловая или нет как у тебя стоит с галочкой так и оставим
//Описание

type Flat struct {
	gorm.Model
	FlatID      int `gorm:"AUTO_INCREMENT"`
	Area        string
	LandMark    string
	Address     string
	HomeNumber  int
	FlatNumber  int
	PriceMin    int
	PriceMax    int
	RoomsCount  int
	RoomsType   sql.NullString
	Floor       int
	FloorsCount int
	Square      int
	FlatType    string
	State       string
	Toilet      string
	ToiletCount int
	BuildYear   int
	IsCorner    bool
	Description string
	//PictureURLs  []string
	//PhoneNumbers map[string]string
}

func NewFlat(Area string,
	LandMark string,
	Address string,
	HomeNumberStr string,
	FlatNumberStr string,
	PriceMinStr string,
	PriceMaxStr string,
	RoomsCountStr string,
	RoomsType string,
	FloorStr string,
	FloorsCountStr string,
	SquareStr string,
	FlatType string,
	State string,
	Toilet string,
	ToiletCountStr string,
	BuildYearStr string,
	IsCorner bool,
	Description string,
	PictureURLs []string,
	PhoneNumbers map[string]string) (Flat, error) {

	intsMap, err := getInts(
		map[string]string{
			"HomeNumber":  HomeNumberStr,
			"FlatNumber":  FlatNumberStr,
			"PriceMin":    PriceMinStr,
			"PriceMax":    PriceMaxStr,
			"RoomsCount":  RoomsCountStr,
			"Floor":       FloorStr,
			"FloorsCount": FloorsCountStr,
			"Square":      SquareStr,
			"ToiletCount": ToiletCountStr,
			"BuildYear":   BuildYearStr,
		})
	if err != nil {
		return Flat{}, err
	}

	roomsType := sql.NullString{}
	roomsType.String = RoomsType
	roomsType.Valid = RoomsType != ""

	return Flat{
		Area:        Area,
		LandMark:    LandMark,
		Address:     Address,
		HomeNumber:  intsMap["HomeNumber"],
		FlatNumber:  intsMap["FlatNumber"],
		PriceMin:    intsMap["PriceMin"],
		PriceMax:    intsMap["PriceMax"],
		RoomsCount:  intsMap["RoomsCount"],
		RoomsType:   roomsType,
		Floor:       intsMap["Floor"],
		FloorsCount: intsMap["FloorsCount"],
		Square:      intsMap["Square"],
		FlatType:    FlatType,
		State:       State,
		Toilet:      Toilet,
		ToiletCount: intsMap["ToiletCount"],
		BuildYear:   intsMap["BuildYear"],
		IsCorner:    IsCorner,
		Description: Description,
		//PictureURLs:  PictureURLs,
		//PhoneNumbers: PhoneNumbers,
	}, nil
}

func getInts(mp map[string]string) (map[string]int, error) {
	result := make(map[string]int)
	for key, val := range mp {
		strInt, err := strconv.Atoi(val)

		if err != nil {
			return nil, errors.Wrap(err, "error getting "+key)
		}
		result[key] = strInt
	}

	return result, nil
}
