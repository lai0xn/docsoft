package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/docsoft/internal/services"
	"github.com/lai0xn/docsoft/internal/types"
	uuid "github.com/satori/go.uuid"
)

type WorkSpace struct {
	service services.Workspace
}

func (c *WorkSpace) CreateWorkspace(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value("userID").(uuid.UUID)
	var body types.WorkspacePayload
	json.NewDecoder(r.Body).Decode(&body)
	err := c.service.Create(body.Name, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "Workspace created",
	})
}

func (c *WorkSpace) DeleteWorkspace(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	userID := r.Context().Value("userID").(string)
	err := c.service.Delete(userID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "workspace deleted",
	})
}

func (c *WorkSpace) JoinWorkspace(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	type payload struct {
		Code string `json:"code"`
	}
	var body payload
	json.NewDecoder(r.Body).Decode(&body)
	if err := c.service.JoinWorkspace(body.Code, userID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "workspace joined",
	})
}

func (c *WorkSpace) GetWorkspaces(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	id, err := uuid.FromString(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	workspaces := c.service.GetWorkSpaces(id)
	json.NewEncoder(w).Encode(workspaces)
}

func (c *WorkSpace) LeaveWorkspace(w http.ResponseWriter, r *http.Request) {
	userID := r.Context().Value("userID").(string)
	workspaceID := chi.URLParam(r, "id")
	type payload struct {
		Code string `json:"code"`
	}
	var body payload
	json.NewDecoder(r.Body).Decode(&body)
	if err := c.service.LeaveWorkspace(userID, workspaceID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "workspace joined",
	})
}
