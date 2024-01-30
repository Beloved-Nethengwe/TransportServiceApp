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
}
type Driver struct {
	IDNumber  string `gorm:"primaryKey"`
	Name      string
	Surname   string
	Number    string
	CreatedAt time.Time
}

type Address struct {
	ID        string `gorm:"primaryKey;autoIncrement"`
	Street    string
	City      string
	CreatedAt time.Time
	ParentID  string
	DriverID  string
	Parent    Parent `gorm:"foreignkey:ParentID;references:IDNumber"`
	Driver    Driver `gorm:"foreignkey:DriverID;references:IDNumber"`
}

type Child struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	Name       string
	Surname    string
	SchoolName string
	State      bool
	ParentID   string
	DriverID   string
	Parent     Parent `gorm:"foreignkey:ParentID;references:IDNumber"`
	Driver     Driver `gorm:"foreignkey:DriverID;references:IDNumber"`
}

type Destination struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	SchoolName string
	DriverID   string
	Driver     Driver `gorm:"foreignkey:DriverID;references:IDNumber"`
}
