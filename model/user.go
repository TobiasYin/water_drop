package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	OpenID string
	Nickname string
	Avatar string
	Gender bool
}
