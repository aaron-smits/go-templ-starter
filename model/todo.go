package model

type Todo struct {
	ID     string `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Done   bool   `json:"done"`
}

type CreateTodoRequest struct {
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Done   bool   `json:"done"`
}

type UpdateTodoRequest struct {
	ID     string `json:"id"`
	UserID int    `json:"user_id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Done   bool   `json:"done"`
}

type DeleteTodoRequest struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type GetTodoRequest struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
}

type GetTodosRequest struct {
	UserID int `json:"user_id"`
}

type GetTodosResponse struct {
	Todos []Todo `json:"todos"`
}

type GetTodoResponse struct {
	Todo Todo `json:"todo"`
}

type CreateTodoResponse struct {
	Todo Todo `json:"todo"`
}

type UpdateTodoResponse struct {
	Todo Todo `json:"todo"`
}

type DeleteTodoResponse struct {
	Todo Todo `json:"todo"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
