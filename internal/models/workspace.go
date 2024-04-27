package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Workspace struct {
	gorm.Model
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Owner   User
	OwnerID uuid.UUID
	Users   []User `gorm:"many2many:workspace_users"`
	Roles   []Role
	Code    string
	Root    *Folder
	RootID  uuid.UUID
}

func (e *Workspace) BeforeCreate(tx *gorm.DB) error {
	e.ID = uuid.NewV4()
	return nil
}
