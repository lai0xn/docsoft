package controllers

import (
	"net/http"

	"github.com/lai0xn/docsoft/internal/services"
)

type Capsule struct {
	service services.Capsule
}

func (c *Capsule) Create(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create capsule"))
}

func (c *Capsule) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get capsule by id"))
}

func (c *Capsule) GetByWorkspace(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get  capsule by workspace"))
}

func (c *Capsule) Revert(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("revert a workspace from a  capsule"))
}

func (c *Capsule) Delete(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete capsule"))
}
