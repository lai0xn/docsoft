package services

import (
	"os"

	"github.com/lai0xn/docsoft/internal/models"
	"github.com/lai0xn/docsoft/storage/db"
	uuid "github.com/satori/go.uuid"
)

type File struct{}

func (f *File) CreateFile(filename string, description string, parentid string) {
	file := models.File{
		File_name:        filename,
		File_description: description,
		File:             "/uploads" + filename,
	}
	db.DB.Create(&file)
}

func (f *File) DeleteFile(id uuid.UUID) error {
	var file models.File
	if err := db.DB.Where("id = ?", id).Find(&file).Error; err != nil {
		return err
	}
	os.Remove(file.File)
	db.DB.Delete(&file)
	return nil
}

func (f *File) RenameFile(id uuid.UUID, new_name string) error {
	var file models.File
	if err := db.DB.Where("id = ?", id).Find(id).Error; err != nil {
		return err
	}
	os.Rename(file.File, new_name)
	file.File_name = new_name
	db.DB.Save(&file)
	return nil
}

func (f *File) MoveFile(id uuid.UUID, new_parent uuid.UUID) error {
	var file models.File
	if err := db.DB.Where("id = ?", id).Find(&file).Error; err != nil {
		return err
	}
	file.FolderID = &new_parent
	db.DB.Save(&file)
	return nil
}
