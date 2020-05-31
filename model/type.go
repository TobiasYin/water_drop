package model

import "github.com/jinzhu/gorm"

type Type struct {
	gorm.Model
	Name string
}
