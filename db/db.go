package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/aaron-smits/templ-starter/model"
	_ "github.com/lib/pq"
	supa "github.com/nedpals/supabase-go"
)

// DB is a wrapper around the database connection.
type DB interface {
	// CreateTodoTable creates the todo table in the database.
	CreateTodoTable() error
	// GetTodoList returns a list of todos.
	GetTodoList() ([]model.Todo, error)
	// AddTodo adds a todo to the database.
	AddTodo(model.Todo) error
	// DeleteTodo deletes a todo from the database.
	// DeleteTodo() error
	// // UpdateTodo updates a todo in the database.
	// UpdateTodo() error
}

type PostgresDB struct {
	DB *sql.DB
}

type SupabaseDB struct {
	DB *supa.Client
}

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
	_, err := DB.DB.Exec("CREATE TABLE IF NOT EXISTS todos (id SERIAL PRIMARY KEY, title TEXT, body TEXT, done BOOLEAN, user_id TEXT)")
	if err != nil {
		return err
	}
	return nil
}

func (DB *PostgresDB) GetTodoList() ([]model.Todo, error) {
	rows, err := DB.DB.Query("SELECT * FROM todos")
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
	fmt.Println("This is the todo: ", todo)
	_, err := DB.DB.Exec("INSERT INTO todos (title, body, done, user_id) VALUES ($1, $2, $3, $4)", todo.Title, todo.Body, todo.Done, todo.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (DB *PostgresDB) DeleteTodo(id int) error {
	_, err := DB.DB.Exec("DELETE FROM todos WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

func (DB *PostgresDB) UpdateTodo(todo model.Todo) error {
	_, err := DB.DB.Exec("UPDATE todos SET title = $1, body = $2, done = $3 WHERE id = $4", todo.Title, todo.Body, todo.Done, todo.ID)
	if err != nil {
		return err
	}
	return nil
}
