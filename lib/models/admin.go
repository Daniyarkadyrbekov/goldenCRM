package models

import "github.com/jinzhu/gorm"

type Address struct {
	gorm.Model
	Address string
}

type Landmark struct {
	gorm.Model
	Landmark string
}
