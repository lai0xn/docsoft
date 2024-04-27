package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"primary_key;type:uuid"`
	First_Name   string
	Last_Name    string
	Email        string `gorm:"unique"`
	Phone_Number string
	Company_Name string
	Password     string
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	u.ID = uuid.NewV4()
	return nil
}
