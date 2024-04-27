package services

import (
	"github.com/lai0xn/docsoft/internal/models"
	"github.com/lai0xn/docsoft/storage/db"
	uuid "github.com/satori/go.uuid"
)

type Folder struct{}

func (f *Folder) CreateFolder(name string, parentID string) error {
	parent, err := uuid.FromString(parentID)
	if err != nil {
		return err
	}
	folder := models.Folder{
		Name:     name,
		ParentID: parent,
	}
	if err := db.DB.Create(&folder).Error; err != nil {
		return err
	}
	return nil
}

func (f *Folder) DeleteFolder(id uuid.UUID) error {
	var folder models.Folder
	db.DB.Where("id = ", id).Find(&folder)
	if err := db.DB.Delete(&folder).Error; err != nil {
		return err
	}
	return nil
}

func (f *Folder) MoveFolder(id uuid.UUID, dest uuid.UUID) error {
	var folder models.Folder
	db.DB.Where("id = ?", id).Find(&folder)
	folder.ParentID = dest
	db.DB.Save(&folder)
	return nil
}

func (f *Folder) RenameFolder(id uuid.UUID, name string) error {
	var folder models.Folder
	db.DB.Where("id = ?", id).Find(&folder)
	folder.Name = name
	db.DB.Save(&folder)
	return nil
}

func (f *Folder) GetChildren(folder_id string) ([]Folder, []File) {
	folder_uuid, err := uuid.FromString(folder_id)
	var folders []Folder
	var files []File
	if err != nil {
		return nil, nil
	}
	db.DB.Where("parent_id = ?", folder_uuid).Find(&folders)
	db.DB.Where("folder_id = ?", folder_uuid).Find(&files)
	return folders, files
}
