package models

import (
	"time"
)

type Parent struct {
	ID              string `gorm:"primaryKey"`
	IDNumber        string
	PName           string `gorm:"alias:parentName"`
	Surname         string
	CellphoneNumber string
	Address         string
	CreatedAt       time.Time
	RoleID          int
	Role            Role `gorm:"foreignkey:RoleID;references:ID"`
}
type Child struct {
	ID           int    `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"alias:childName"`
	Surname      string
	Allergy      string
	EmergContact string
	PickUp       string
	Destination  string
	//RelationShips
	ParentID string
	Parent   Parent `gorm:"foreignkey:ParentID;references:ID"`
}

type Role struct {
	ID   int `gorm:"primaryKey;autoIncrement"`
	Role string
}
type Driver struct {
	ID                    string `gorm:"primaryKey"`
	IDNumber              string
	Name                  string
	Surname               string
	CellphoneNumber       string
	Image                 string
	CarRegistrationNumber string
	CreatedAt             time.Time
	RoleID                int
	Role                  Role `gorm:"foreignkey:RoleID;references:ID"`
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
	ParentID string
	Parent   Parent `gorm:"foreignkey:ParentID;references:ID"`
	DriverID string
	Driver   Driver `gorm:"foreignkey:DriverID;references:ID"`
	ChildID  int
	Child    Child `gorm:"foreignkey:ChildID;references:ID"`
}
