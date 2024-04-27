package models

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type TimeCapsule struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
	WorkspaceID uuid.UUID
	Workspace   Workspace
	CreatedBy   User
	CreatedByID uuid.UUID
}

func (t *TimeCapsule) BeforeCreate(tx *gorm.DB) error {
	t.ID = uuid.NewV4()
	return nil
}
