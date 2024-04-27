package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key"`
	Name        string
	Workspace   Workspace
	WorkspaceID uuid.UUID
	Owners      []User `gorm:"many2many:workspace_role"`
}

func (r *Role) BeforeCreate(tx *gorm.DB) error {
	r.ID = uuid.NewV4()
	return nil
}
