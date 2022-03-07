package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID    uint64 `gorm:"primary_key:auto_increment"`
	Name  string
	Email string
	Notes []*Note `json:"items"`
}

type Note struct {
	gorm.Model           // Adds some metadata fields to the table
	ID         uuid.UUID `gorm:"type:uuid"` // Explicitly specify the type to be uuid
	Title      string
	SubTitle   string
	Text       string
	UserID     uint64 `gorm:"not null" json"-"`
	User       User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,ondelete:CASCADE" json:"user"`
}
