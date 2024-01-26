package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/aaron-smits/templ-starter/model"
	_ "github.com/lib/pq"
	supa "github.com/nedpals/supabase-go"
)

var TodoList []model.Todo

// DB is a wrapper around the database connection.
type DB interface {
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
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	return &PostgresDB{DB: db}, nil
}

func NewSupabaseDB() (*SupabaseDB, error) {
	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")
	supabase := supa.CreateClient(supabaseURL, supabaseKey)
	return &SupabaseDB{DB: supabase}, nil
}
