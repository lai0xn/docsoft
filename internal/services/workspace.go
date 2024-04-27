package services

import (
	"errors"

	"github.com/lai0xn/docsoft/internal/models"
	"github.com/lai0xn/docsoft/storage/db"
	"github.com/lai0xn/docsoft/utils"
	uuid "github.com/satori/go.uuid"
)

type Workspace struct{}

func (s *Workspace) Create(name string, owner_id uuid.UUID) error {
	// we generate the id first to avoid the conflict of having 2 different ids generated
	id := uuid.NewV4()

	parent_folder := models.Folder{
		Name:        "root",
		WorkspaceID: id,
	}
	db.DB.Create(parent_folder)
	workspace := models.Workspace{
		OwnerID: owner_id,
		ID:      id,
		RootID:  parent_folder.ID,
		Code:    utils.GenerateCode(),
	}
	db.DB.Create(&workspace)
	return nil
}

func (s *Workspace) Delete(userID string, id string) error {
	workspace_uuid, err := uuid.FromString(id)
	if err != nil {
		return err
	}
	user_uuid, err := uuid.FromString(userID)
	if err != nil {
		return err
	}
	var workspace models.Workspace
	db.DB.Where("id = ?", workspace_uuid).Find(&workspace)
	if workspace.OwnerID != user_uuid {
		return errors.New("You don't have the permissions to do taht")
	}
	db.DB.Delete(&workspace)
	return nil
}

func (s *Workspace) JoinWorkspace(code string, userID string) error {
	var workspace models.Workspace
	var user models.User
	if err := db.DB.Where("code = ?", code).Find(&workspace).Error; err != nil {
		return err
	}
	db.DB.Model(&workspace).Association("Users").Append(user)
	return nil
}

func (s *Workspace) GetWorkSpaces(userID uuid.UUID) []models.Workspace {
	var user models.User
	var workspaces []models.Workspace
	if err := db.DB.Where("id = ?", userID).Find(&user).Error; err != nil {
		return nil
	}
	if err := db.DB.Joins("JOIN workspace_users ON workspaces.id = workspace_users.workspace_id").
		Where("workspace_users.user_id = ?", userID).
		Find(&workspaces).Error; err != nil {
		return nil
	}
	return workspaces
}

func (s *Workspace) LeaveWorkspace(userID string, workspaceID string) error {
	var workspace models.Workspace
	var user models.User
	user_uuid, err := uuid.FromString(userID)
	if err != nil {
		return err
	}
	workspace_uuid, err := uuid.FromString(workspaceID)
	if err != nil {
		return err
	}
	err = db.DB.Where("id = ?", user_uuid).Find(&user).Error
	if err != nil {
		return err
	}
	err = db.DB.Where("id = ?", workspace_uuid).Find(&workspace).Error
	if err != nil {
		return err
	}
	err = db.DB.Model(&workspace).Association("Users").Delete(user)
	if err != nil {
		return err
	}
	return nil
}

func (s *Workspace) GetByID(id uuid.UUID) (models.Workspace, error) {
	var workspace models.Workspace
	err := db.DB.Where("id = ?", id).Find(&workspace).Error
	if err != nil {
		return models.Workspace{}, err
	}
	return workspace, nil
}
