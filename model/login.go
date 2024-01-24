package model

// these do not currently have any use in the project,
// but we can adapt them to our needs later
// These would be good to have for type safety across the project

type LoginRequest struct {
	AccountNumber int64  `json:"account_number"`
	Password      string `json:"password"`
}

type LoginResponse struct {
	AccountNumber int64  `json:"sub"` // Sub is part of the JWT spec
	Token         string `json:"access_token"`
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Balance   int64  `json:"balance"`
	IsAdmin   bool   `json:"is_admin"`
}

type UpdateAccountRequest struct {
	FirstName     string `json:"first_name"`
	LastName      string `json:"last_name"`
	AccountNumber int64  `json:"account_number"`
	IsAdmin       bool   `json:"is_admin"`
}
