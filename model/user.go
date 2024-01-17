package model

import "time"

type User struct {
	ID   int `json:"id"`
	OrgID int `json:"org_id"`
	Email string `json:"email"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	EncryptedPassword string `json:"encrypted_password"`
	IsAdmin bool `json:"is_admin"`
	IsArchived bool `json:"is_archived"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
