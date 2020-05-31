package model

import (
	"github.com/jinzhu/gorm"
	"time"
)

const (
	ImportanceL1 = iota
	ImportanceL2
	ImportanceL3
	ImportanceL4
	ImportanceL5
)

const (
	RepeatL1 = iota
	RepeatL2
	RepeatL3
	RepeatL4
	RepeatL5
)

const (
	StatusOK = iota
	StatusFail
	StatusCancel
)

type Card struct {
	gorm.Model
	Title        string
	Description  string `gorm:"type:text;"`
	Importance   int
	Deadline     time.Time
	Repeat       int
	Status       int
	Image        string
	Diary        string `gorm: "type:text;"`
	Progress     uint8
	Style        string `gorm: "type:text;"`
	Owner        User
	Participants []User `gorm:"many2many:card_participants;"`
	CardSet      CardSet
	Type         []Type `gorm:"many2many:card_type;"`
}
