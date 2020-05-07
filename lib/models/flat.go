package models

import "math/rand"

type Flat struct {
	id          int64
	street      string
	home        string
	structure   uint
	flatNumber  uint
	state       flatState
	floor       uint
	isCorner    bool
	flatType    string
	description string
	pictureURLs []string
	owner       string
}

func NewFlat(street string,
	home string,
	structure uint,
	flatNumber uint,
	state flatState,
	floor uint,
	isCorner bool,
	flatType string,
	description string,
	pictureURLs []string,
	owner string) Flat {

	return Flat{
		id:          rand.Int63(),
		street:      street,
		home:        home,
		structure:   structure,
		flatNumber:  flatNumber,
		state:       state,
		floor:       floor,
		isCorner:    isCorner,
		flatType:    flatType,
		description: description,
		pictureURLs: pictureURLs,
		owner:       owner,
	}
}

type flatState string

const (
	Good   flatState = "хорошее"
	Euro   flatState = "евро"
	Soviet flatState = "советское"
)
