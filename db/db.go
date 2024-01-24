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
	CreateUser(*model.User) (*model.User, error)
	DeleteUser(int) error
	UpdateUserByID(int, *model.User) (*model.User, error)
	GetUsers() ([]*model.User, error)
	GetUserByID(int) (*model.User, error)
	GetUserByNumber(string) (*model.User, error)
	GetAdminStatus(int) (bool, error)
	MakeTransfer(int, int, int) (*model.User, error)
	AddBalanceTx(*sql.Tx, int, int) error
	SubtractBalanceTx(*sql.Tx, int, int) error
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

// func NewSupabaseDB() (*SupabaseDB, error) {
// 	supabaseURL := os.Getenv("SUPABASE_URL")
// 	supabaseKey := os.Getenv("SUPABASE_KEY")
// 	supabase := supa.CreateClient(supabaseURL, supabaseKey)
// 	return &SupabaseDB{DB: supabase}, nil
// }

func (db *PostgresDB) CreateUserTable() error {
	_, err := db.DB.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		is_archived BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("error creating user table: %v", err)
	}
	return nil
}

func (db *PostgresDB) CreateUser(u *model.User) (*model.User, error) {
	query := `INSERT INTO users (email, first_name, last_name, is_archived, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6);`
	row := db.DB.QueryRow(
		query,
		u.Email, u.FirstName, u.LastName, u.IsArchived, u.CreatedAt, u.UpdatedAt)

	err := row.Scan(&u.ID)
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %v", err)
	}
	return u, nil
}

func (db *PostgresDB) DeleteUser(id int) error {
	// set the is_archived column to true
	_, err := db.DB.Exec(`UPDATE users SET is_archived = true WHERE id = $1;`, id)
	if err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}

	return nil
}

func (db *PostgresDB) UpdateUserByID(id int, u *model.User) (*model.User, error) {
	query := `
	UPDATE users
	SET email = $1, first_name = $2, last_name = $3, is_archived = $4, updated_at = $5`
	if u == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	_, err := db.DB.Exec(query, u.Email, u.FirstName, u.LastName, u.IsArchived, u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	user, err := db.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (db *PostgresDB) GetUserByID(id int) (*model.User, error) {
	user := new(model.User)
	query := `SELECT * FROM users WHERE id = $1;`
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.IsArchived, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %v", err)
	}

	return user, nil
}
