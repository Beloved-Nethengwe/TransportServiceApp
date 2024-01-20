package models

import (
	"time"
)

type Parent struct {
	ID        string `gorm:"primaryKey"`
	Name      string
	Surname   string
	Number    string
	CreatedAt time.Time
	Address   Address `gorm:"foreignkey:ParentIDID"`
}

type Address struct {
	ID        uint `gorm:"primary_key"`
	City      string
	ParentID  uint
	CreatedAt time.Time
}
