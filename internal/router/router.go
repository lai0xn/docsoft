package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/lai0xn/docsoft/internal/controllers"
	"github.com/lai0xn/docsoft/internal/middlewares"
)

func Setup(r chi.Router) {
	// auth routes
	r.Group(func(r chi.Router) {
		auth := controllers.Auth{}
		r.Post("/auth/login/", auth.Login)
		r.Post("/auth/signup/", auth.Signup)
	})

	// files routes
	r.Group(func(r chi.Router) {
		f := controllers.File{}
		r.Use(middlewares.LoginRequired)
		r.Post("/files/upload/", f.UploadFile)
		r.Delete("/files/delete/{id}", f.DeleteFile)
		r.Post("/files/rename/{id}", f.RenameFile)
		r.Post("/files/move/{id}", f.MoveFile)
	})
	// folders routes
	r.Group(func(r chi.Router) {
		f := controllers.Folder{}
		r.Use(middlewares.LoginRequired)
		r.Post("/folders/create/", f.CreateFolder)
		r.Delete("/folders/delete/{id}", f.DeleteFolder)
		r.Post("/folders/rename/{id}", f.RenameFolder)
		r.Post("/folders/move/{id}", f.MoveFolder)
		r.Get("/folders/folder/{id}", f.GetChildren)
	})

	// workspaces routes
	r.Group(func(r chi.Router) {
		workspace := controllers.WorkSpace{}
		r.Use(middlewares.LoginRequired)
		r.Post("/workspaces/create/", workspace.CreateWorkspace)
		r.Post("/workspaces/delete/{id}", workspace.DeleteWorkspace)
		r.Get("/workspaces/me", workspace.GetWorkspaces)
		r.Post("/workspaces/join", workspace.JoinWorkspace)
		r.Post("/workspaces/leave/{id}", workspace.LeaveWorkspace)
	})
	r.Group(func(r chi.Router) {
		roles := controllers.Roles{}
		r.Use(middlewares.LoginRequired)
		r.Post("/roles/create/{id}", roles.CreateRole)
		r.Post("/roles/{role_id}/assign/{user_id}", roles.AssignRole)
		r.Get("/roles/role/{role_id}", roles.GetRole)
		r.Delete("/roles/role/{role_id}", roles.DeleteRole)
	})
	r.Group(func(r chi.Router) {
		capsules := controllers.Capsule{}
		r.Use(middlewares.LoginRequired)
		r.Post("/capsules/create/", capsules.Create)
		r.Post("/capsules/{capsule_id}/revert/{workspace_id}", capsules.Revert)
		r.Get("/capsules/workspace/{id}", capsules.GetByWorkspace)
		r.Delete("/capsules/delete/{id}", capsules.Delete)
	})
}
