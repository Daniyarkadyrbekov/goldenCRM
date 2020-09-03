package models

import "github.com/jinzhu/gorm"

type Owner struct {
	gorm.Model
	Name    string
	Phone   string
	OwnerID int
}
