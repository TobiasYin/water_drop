package model

import "github.com/jinzhu/gorm"

const (
	TypeCard = iota
	TypeCardSet
)

type Market struct {
	gorm.Model
	Type     int
	EntityID int
}
