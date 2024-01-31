package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/aaron-smits/templ-starter/model"
	_ "github.com/lib/pq"
	supa "github.com/nedpals/supabase-go"
)

// DB is a wrapper around the database connection.
// Currently, you can create, read, and delete todos.
type DB interface {
	// CreateTodoTable creates the todo table in the database.
	CreateTodoTable() error
	// GetTodoList returns a list of todos.
	GetTodoList() ([]model.Todo, error)
	// AddTodo adds a todo to the database.
	AddTodo(model.Todo) error
	// DeleteTodo deletes a todo from the database.
	DeleteTodo(id string) error
}

// type for postgres client. This is currently used for the DB handling
type PostgresDB struct {
	DB *sql.DB
}

// type for supabase client in case we want to use it for SQL queries
type SupabaseDB struct {
	DB *supa.Client
}

// NewPostgresDB creates a new PostgresDB instance.
func NewPostgresDB() (*PostgresDB, error) {
	connectionString := os.Getenv("POSTGRES_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("POSTGRES_CONNECTION_STRING environment variable must be set")
	}
	DB, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return &PostgresDB{DB: DB}, nil
}

func NewSupabaseDB() (*SupabaseDB, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)
	return &SupabaseDB{DB: supabase}, nil
}

func (DB *PostgresDB) CreateTodoTable() error {
	_, err := DB.DB.Query("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, title TEXT, body TEXT, done BOOLEAN, user_id TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (DB *PostgresDB) GetTodoList() ([]model.Todo, error) {
	query := "SELECT * FROM todos"
	rows, err := DB.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []model.Todo
	for rows.Next() {
		t := model.Todo{}
		if err := rows.Scan(&t.ID, &t.Title, &t.Body, &t.Done, &t.UserID); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

func (DB *PostgresDB) AddTodo(todo model.Todo) error {
	query := "INSERT INTO todos (title, body, done, user_id) VALUES ($1, $2, $3, $4)"
	_, err := DB.DB.Query(query, todo.Title, todo.Body, todo.Done, todo.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (DB *PostgresDB) DeleteTodo(id string) error {
	query := "DELETE FROM todos WHERE id = $1"
	_, err := DB.DB.Query(query, id)
	if err != nil {
		return err
	}
	return nil
}

// func (DB *PostgresDB) UpdateTodo(todo model.Todo) error {
// 	query := "UPDATE todos SET title = $1, body = $2, done = $3 WHERE id = $4"
// 	_, err := DB.DB.Query(query, todo.Title, todo.Body, todo.Done, todo.ID)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
