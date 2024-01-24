package model

import "time"

// these do not currently have any use in the project,
// but we can adapt them to our needs later
// These would be good to have for type safety across the project

type User struct {
	ID         int       `json:"id"`
	OrgID      int       `json:"org_id"`
	Email      string    `json:"email"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	IsArchived bool      `json:"is_archived"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
