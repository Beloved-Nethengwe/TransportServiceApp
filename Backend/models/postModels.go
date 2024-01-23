package models

import (
	"time"
)

type Parent struct {
	IDNumber  string `gorm:"primaryKey"`
	Name      string
	Surname   string
	Number    string
	CreatedAt time.Time
	Address   Address `gorm:"foreignkey:ParentID"`
}

type Address struct {
	ID        string `gorm:"primary_key"`
	Street    string
	City      string
	ParentID  string
	CreatedAt time.Time
}
