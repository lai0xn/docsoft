package services

import (
	"github.com/lai0xn/docsoft/internal/models"
	"github.com/lai0xn/docsoft/internal/types"
	"github.com/lai0xn/docsoft/storage/db"
	uuid "github.com/satori/go.uuid"
)

type Roles struct{}

func (s *Roles) CreateRole(payload types.RolePayload) error {
	workspaceID, err := uuid.FromString(payload.WorkspaceID)
	if err != nil {
		return err
	}
	role := models.Role{
		Name:        payload.Name,
		WorkspaceID: workspaceID,
	}
	db.DB.Create(&role)
	return nil
}

func (s *Roles) DeleteRole(role_id string) error {
	var role models.Role
	id, err := uuid.FromString(role_id)
	if err != nil {
		return err
	}
	if err := db.DB.Where("id = ?", id).Delete(&role).Error; err != nil {
		return err
	}
	return nil
}

func (s *Roles) GetRole(role_id uuid.UUID) (models.Role, error) {
	var role models.Role
	err := db.DB.Where("id = ?", role_id).Preload("Workspace").Find(role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}

func (s *Roles) AssignRole(role_id uuid.UUID, user_id uuid.UUID) error {
	var user models.User
	var role models.Role

	if err := db.DB.Where("id = ?", role_id).Find(&role).Error; err != nil {
		return err
	}

	if err := db.DB.Where("id = ?", user_id).Find(&user).Error; err != nil {
		return err
	}

	db.DB.Model(&role).Association("Owners").Append(&user)
	return nil
}

func (s *Roles) WorkspaceRoles(workspace_id string) []models.Role {
	var roles []models.Role
	workspace_uuid, err := uuid.FromString(workspace_id)
	if err != nil {
		return nil
	}
	if err := db.DB.Where("workspace_id = ?", workspace_uuid).Find(&roles).Error; err != nil {
		return nil
	}
	return roles
}
