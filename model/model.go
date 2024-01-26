package model

import (
	supa "github.com/nedpals/supabase-go"
)

type User struct {
	*supa.User
}

type Todo struct {
	ID     int    `json:"id"`
	UserID string `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Done   bool   `json:"done"`
}

type TodoList struct {
	Todos []Todo `json:"todos"`
}
