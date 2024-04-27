package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/docsoft/internal/services"
	"github.com/lai0xn/docsoft/internal/types"
	uuid "github.com/satori/go.uuid"
)

type Roles struct {
	service           services.Roles
	workspace_service services.Workspace
}

func (c *Roles) CreateRole(w http.ResponseWriter, r *http.Request) {
	var paylaod types.RolePayload
	id := chi.URLParam(r, "id")
	workspace_uuid, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("userID").(string)
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	workspace, err := c.workspace_service.GetByID(workspace_uuid)
	if workspace.OwnerID != userUUID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if id != userID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	json.NewDecoder(r.Body).Decode(&paylaod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.CreateRole(paylaod)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "role created",
	})
}

func (c *Roles) AssignRole(w http.ResponseWriter, r *http.Request) {
	currentID := r.Context().Value("userID").(string)
	currentUUID, err := uuid.FromString(currentID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	role_id := chi.URLParam(r, "role_id")
	roleUUID, err := uuid.FromString(role_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	user_id := chi.URLParam(r, "user_id")
	userUUID, err := uuid.FromString(user_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	role, err := c.service.GetRole(roleUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if role.Workspace.OwnerID != currentUUID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	err = c.service.AssignRole(roleUUID, userUUID)
}

func (c *Roles) DeleteRole(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	workspace_uuid, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	userID := r.Context().Value("userID").(string)
	userUUID, err := uuid.FromString(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	workspace, err := c.workspace_service.GetByID(workspace_uuid)
	if workspace.OwnerID != userUUID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if id != userID {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = c.service.DeleteRole(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(types.H{
		"message": "role created",
	})
}

func (c *Roles) GetRole(w http.ResponseWriter, r *http.Request) {
	role_id := chi.URLParam(r, "role_id")
	roleUUID, err := uuid.FromString(role_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	role, err := c.service.GetRole(roleUUID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	json.NewEncoder(w).Encode(role)
}
