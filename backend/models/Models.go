package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserID      string
	Email       string `gorm:"uniqueIndex;not null"`
	Password    string `gorm:"not null"`
	PhoneNumber string
	FirstName   string
	LastName    string
}

type Satelitte struct {
	gorm.Model
	SatelitteID       string
	SatelitteName     string
	SatelitteLocation string
	ContactPerson     string
	ContactEmail      string
	ContactPassword   string
}

type Donor struct {

    UserID       string
    DonationDate string
    BloodType    string
    Status       string
    SatelliteID  string
	RegionalID string
    SourceType   string // New field to track the source
}

type Regional struct {
	ID string
	RegionID       string `gorm:"not null"`
	RegionName     string `gorm:"not null"`
	RegionLocation string `gorm:"not null"`
	ContactPerson     string `gorm:"not null"`
	ContactEmail      string `gorm:"not null"`
	ContactPassword   string `gorm:"not null"`
}
