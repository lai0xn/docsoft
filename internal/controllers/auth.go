package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lai0xn/docsoft/internal/services"
	"github.com/lai0xn/docsoft/internal/types"
)

type Auth struct {
	service *services.Auth
}

func (c *Auth) Login(w http.ResponseWriter, r *http.Request) {
	var payload types.LoginPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	token, err := c.service.GenerateToken(payload.Email, payload.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(types.H{
		"token": token,
	})
}

func (c *Auth) Signup(w http.ResponseWriter, r *http.Request) {
	var payload types.SignupPayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	err = c.service.Signup(payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(types.H{
		"message": "user crreated",
	})
}
