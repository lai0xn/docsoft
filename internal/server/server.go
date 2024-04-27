package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/lai0xn/docsoft/internal/router"
	"github.com/lai0xn/docsoft/storage/db"
)

type Server struct {
	ADDR string
	Mux  chi.Router
}

func (s *Server) Setup() {
	db.Connect()
	s.Mux = chi.NewRouter()
	s.Mux.Use(middleware.Logger)
	// Configure CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // Allow requests from any origin
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		MaxAge:         300, // Maximum age for preflight requests
	})

	// Use CORS middleware
	s.Mux.Use(corsMiddleware.Handler)
	router.Setup(s.Mux)
}

func (s *Server) Run() {
	s.Setup()
	fmt.Println("Server Running on port", s.ADDR)

	http.ListenAndServe(":"+s.ADDR, s.Mux)
}
