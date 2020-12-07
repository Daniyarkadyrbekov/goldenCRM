package models

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type Flat struct {
	gorm.Model
	FlatID      int `gorm:"AUTO_INCREMENT"`
	Area        string
	LandMark    string
	Address     string
	HomeNumber  int
	Building    string
	FlatNumber  int
	PriceMin    int
	PriceMax    int
	RoomsCount  int
	RoomsType   string
	Floor       int
	FloorsCount int
	Square      int
	FlatType    string
	State       string
	Toilet      string
	ToiletCount int
	BuildYear   int
	IsCorner    bool
	IsSeparated bool
	Description string
	Owners      []Owner `gorm:"ForeignKey:OwnerID"`
}

func NewFlat(Area string,
	LandMark string,
	Address string,
	HomeNumberStr string,
	Building string,
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
	isSeparated bool,
	Description string,
	owners []Owner) (Flat, error) {

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

	return Flat{
		Area:        Area,
		LandMark:    LandMark,
		Address:     Address,
		HomeNumber:  intsMap["HomeNumber"],
		Building:    Building,
		FlatNumber:  intsMap["FlatNumber"],
		PriceMin:    intsMap["PriceMin"],
		PriceMax:    intsMap["PriceMax"],
		RoomsCount:  intsMap["RoomsCount"],
		RoomsType:   RoomsType,
		Floor:       intsMap["Floor"],
		FloorsCount: intsMap["FloorsCount"],
		Square:      intsMap["Square"],
		FlatType:    FlatType,
		State:       State,
		Toilet:      Toilet,
		ToiletCount: intsMap["ToiletCount"],
		BuildYear:   intsMap["BuildYear"],
		IsCorner:    IsCorner,
		IsSeparated: isSeparated,
		Description: Description,
		Owners:      owners,
	}, nil
}

func getInts(mp map[string]string) (map[string]int, error) {
	result := make(map[string]int)
	for key, val := range mp {
		if val == "" {
			continue
		}

		strInt, err := strconv.Atoi(val)

		if err != nil {
			return nil, errors.Wrap(err, "error getting "+key)
		}
		result[key] = strInt
	}

	return result, nil
}
