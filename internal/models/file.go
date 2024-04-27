package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type File struct {
	gorm.Model
	ID               uuid.UUID `gorm:"primary_key;type:uuid"`
	File_name        string
	File_description string
	File             string
	Folder           *Folder
	FolderID         *uuid.UUID
}

func (e *File) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.NewV4()
	return nil
}
