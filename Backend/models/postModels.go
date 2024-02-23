package models

import (
	"time"
)

type Parent struct {
	ID              string `gorm:"primaryKey"`
	IDNumber        string
	Name            string
	Surname         string
	CellphoneNumber string
	Address         string
	Email           string
	Password        string
	CreatedAt       time.Time
}
type Child struct {
	ID           int `gorm:"primaryKey;autoIncrement"`
	Name         string
	Surname      string
	Allergy      string
	EmergContact string
	PickUp       string
	Destination  string
	//RelationShips
	ParentID string
	Parent   Parent `gorm:"foreignkey:ParentID;references:ID"`
}
type Driver struct {
	ID                    int `gorm:"primaryKey;autoIncrement"`
	IDNumber              string
	Name                  string
	Surname               string
	CellphoneNumber       string
	Image                 string
	CarRegistrationNumber string
	Email                 string
	Password              string
	CreatedAt             time.Time
}

type Destination struct {
	ID         int `gorm:"primaryKey;autoIncrement"`
	SchoolName string

	//RelationShips
	DriverID string
	Driver   Driver `gorm:"foreignkey:DriverID;references:ID"`
}

type RequestBridge struct {
	ID     int `gorm:"primaryKey;autoIncrement"`
	Status string

	//RelationShips
	ParentID int
	Parent   Parent `gorm:"foreignkey:ParentID;references:ID"`
	DriverID int
	Driver   Driver `gorm:"foreignkey:DriverID;references:ID"`
	ChildID  int
	Child    Child `gorm:"foreignkey:ChildID;references:ID"`
}
