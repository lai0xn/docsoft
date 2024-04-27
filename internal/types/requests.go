package types

type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupPayload struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	First_Name   string `json:"first_name"`
	Last_Name    string `json:"last_name"`
	Company_Name string `json:"company_name"`
	Phone_Number string `json:"phone_number"`
}

type FolderPayload struct {
	Name     string `json:"name"`
	ParentID string `json:"parent_id"`
}

type WorkspacePayload struct {
	Name string `json:"name"`
}

type RenamePayload struct {
	Name string `json:"string"`
}

type RolePayload struct {
	Name        string `json:"name"`
	WorkspaceID string `json:"workspace_id"`
}

type MovePayload struct {
	ParentID string `json:"parent_id"`
}

type CapsulePaylaod struct {
	WorkspaceID string `json:"workspace_id"`
	CreatedByID string `json:"created_by"`
}

type H map[string]interface{}
