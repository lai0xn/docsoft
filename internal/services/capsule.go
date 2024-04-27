package services

import (
	"github.com/lai0xn/docsoft/internal/models"
	"github.com/lai0xn/docsoft/internal/types"
	"github.com/lai0xn/docsoft/storage/db"
	uuid "github.com/satori/go.uuid"
)

type Capsule struct{}

func (s *Capsule) Create(payload types.CapsulePaylaod) {
	created_by, _ := uuid.FromString(payload.CreatedByID)
	workspace_id, _ := uuid.FromString(payload.WorkspaceID)
	var workspace models.Workspace
	db.DB.Find("id = ?", workspace_id).Find(&workspace)
	workspace_clone := models.Workspace{
		OwnerID: workspace.OwnerID,
		Root:    workspace.Root,
		Roles:   workspace.Roles,
		Users:   workspace.Users,
	}
	db.DB.Create(&workspace_clone)
	capsule := models.TimeCapsule{
		CreatedByID: created_by,
		WorkspaceID: workspace_clone.ID,
	}
	db.DB.Create(capsule)
}

func (s *Capsule) Delete(id uuid.UUID) {
	db.DB.Where("id = ?", id).Delete(&models.TimeCapsule{})
}

func (s *Capsule) GetByID(id uuid.UUID) (models.TimeCapsule, error) {
	var capsule models.TimeCapsule
	if err := db.DB.Where("id = ?", id).Find(&capsule).Error; err != nil {
		return models.TimeCapsule{}, err
	}
	return capsule, nil
}

func (s *Capsule) GetByWorkspace(workspace_id uuid.UUID) ([]models.TimeCapsule, error) {
	var capsules []models.TimeCapsule
	if err := db.DB.Where("workspace_id", workspace_id).Find(&capsules).Error; err != nil {
		return nil, err
	}
	return capsules, nil
}

func (s *Capsule) Revert(workspace_id uuid.UUID, capsule_id uuid.UUID) error {
	var workspace models.Workspace
	var capsule models.TimeCapsule
	if err := db.DB.Where("id = ?", workspace_id).Find(&workspace).Error; err != nil {
		return err
	}
	if err := db.DB.Where("id = ?", capsule_id).Find(&capsule).Error; err != nil {
		return err
	}
	workspace = capsule.Workspace
	db.DB.Save(&workspace)
	return nil
}
