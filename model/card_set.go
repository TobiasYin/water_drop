package model

import "github.com/jinzhu/gorm"

type CardSet struct {
	gorm.Model
	IsPublic bool
	Title    string
	Image    string
	Style    string `gorm:"type:text;"`
	Owner    User
	Type     []Type `gorm:"many2many:card_set_type;"`
}
