package models

import "math/rand"

type Flat struct {
	ID          int64
	Street      string
	Home        string
	Structure   uint
	FlatNumber  uint
	State       flatState
	Floor       uint
	IsCorner    bool
	FlatType    string
	Description string
	PictureURLs []string
	Owner       string
}

func NewFlat(Street string,
	Home string,
	Structure uint,
	FlatNumber uint,
	State flatState,
	Floor uint,
	IsCorner bool,
	FlatType string,
	Description string,
	PictureURLs []string,
	Owner string) Flat {

	return Flat{
		ID:          rand.Int63(),
		Street:      Street,
		Home:        Home,
		Structure:   Structure,
		FlatNumber:  FlatNumber,
		State:       State,
		Floor:       Floor,
		IsCorner:    IsCorner,
		FlatType:    FlatType,
		Description: Description,
		PictureURLs: PictureURLs,
		Owner:       Owner,
	}
}

type flatState string

const (
	Good   flatState = "хорошее"
	Euro   flatState = "евро"
	Soviet flatState = "советское"
)
