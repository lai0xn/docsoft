package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/docsoft/internal/services"
	"github.com/lai0xn/docsoft/internal/types"
	uuid "github.com/satori/go.uuid"
)

type File struct {
	service services.File
}

func (c *File) UploadFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB limit
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	parentID := r.Form.Get("parentID")
	description := r.Form.Get("description")
	if err != nil {
		http.Error(w, "Error retrieving file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a new file on the server to save the uploaded file
	newFile, err := os.Create("./uploads/" + handler.Filename)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	// Copy the uploaded file to the new file on the server
	_, err = io.Copy(newFile, file)
	if err != nil {
		http.Error(w, "Error copying file", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	c.service.CreateFile(handler.Filename, parentID, description)
	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func (c *File) DeleteFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid_id, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	c.service.DeleteFile(uuid_id)
}

func (c *File) RenameFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type payload struct {
		Name string `json:"string"`
	}
	var file_info payload
	json.NewDecoder(r.Body).Decode(&file_info)
	c.service.RenameFile(uuid, file_info.Name)
	json.NewEncoder(w).Encode(&types.H{
		"message": "file renamed",
	})
}

func (c *File) MoveFile(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uuid_, err := uuid.FromString(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	type payload struct {
		DestID string `json:"destination_id"`
	}
	var dest_info payload
	json.NewDecoder(r.Body).Decode(&dest_info)
	dest_id, err := uuid.FromString(dest_info.DestID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.service.MoveFile(uuid_, dest_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
