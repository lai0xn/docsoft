package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/docsoft/internal/services"
	"github.com/lai0xn/docsoft/internal/types"
	uuid "github.com/satori/go.uuid"
)

type Folder struct {
	service *services.Folder
}

func (c *Folder) CreateFolder(w http.ResponseWriter, r *http.Request) {
	var payload types.FolderPayload
	json.NewEncoder(w).Encode(&payload)
	err := c.service.CreateFolder(payload.Name, payload.ParentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "folder created",
	})
}

func (c *Folder) DeleteFolder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid_, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	err = c.service.DeleteFolder(uuid_)
}

func (c *Folder) RenameFolder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid_, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}
	json.NewEncoder(w).Encode(types.H{
		"message": "folder renamed",
	})

	var payload types.RenamePayload
	json.NewDecoder(r.Body).Decode(&payload)
	c.service.RenameFolder(uuid_, payload.Name)
}

func (c *Folder) MoveFolder(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid_, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return

	}

	var payload types.MovePayload
	json.NewDecoder(r.Body).Decode(&payload)
	dest_id, err := uuid.FromString(payload.ParentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.service.MoveFolder(uuid_, dest_id)
	json.NewEncoder(w).Encode(types.H{
		"message": "folder moved",
	})
}

func (c *Folder) GetChildren(w http.ResponseWriter, r *http.Request) {
	folder_id := chi.URLParam(r, "id")
	folders, files := c.service.GetChildren(folder_id)
	json.NewEncoder(w).Encode(types.H{
		"folders": folders,
		"files":   files,
	})
}
