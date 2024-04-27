package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Folder struct {
	gorm.Model
	ID          uuid.UUID `gorm:"primary_key;type:uuid"`
	Name        string
	Parent      *Folder `gorm:"foreignKey:ParentID"`
	ParentID    uuid.UUID
	WorkspaceID uuid.UUID
	Workspace   Workspace
}

func (e *Folder) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.NewV4()
	return nil
}
