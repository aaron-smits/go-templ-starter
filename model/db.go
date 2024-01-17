package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// DB is a wrapper around the database connection.
type DB interface {
	CreateUser(*User) (*User, error)
	DeleteUser(int) error
	UpdateUserByID(int, *User) (*User, error)
	GetUsers() ([]*User, error)
	GetUserByID(int) (*User, error)
	GetUserByNumber(string) (*User, error)
	GetAdminStatus(int) (bool, error)
	MakeTransfer(int, int, int) (*User, error)
	AddBalanceTx(*sql.Tx, int, int) error
	SubtractBalanceTx(*sql.Tx, int, int) error
}


type PostgresDB struct {
	DB *sql.DB
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

	return &PostgresDB{db: db}, nil
}

func (db *PostgresDB) CreateUserTable() error {
	_, err := db.db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL,
		first_name VARCHAR(255) NOT NULL,
		last_name VARCHAR(255) NOT NULL,
		encrypted_password VARCHAR(255) NOT NULL,
		is_admin BOOLEAN NOT NULL,
		is_archived BOOLEAN NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP NOT NULL
	);`)
	if err != nil {
		return fmt.Errorf("error creating user table: %v", err)
	}
	return nil
}

func (db *PostgresDB) CreateUser(u *User) (*User, error) {
	_, err := db.DB.Exec(`INSERT INTO users (email, first_name, last_name, encrypted_password, is_admin, is_archived, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`, u.Email, u.FirstName, u.LastName, u.EncryptedPassword, u.IsAdmin, u.IsArchived, u.CreatedAt, u.UpdatedAt)
	row := db.db.QueryRow(`SELECT id, email, first_name, last_name, encrypted_password, is_admin, is_archived, created_at, updated_at FROM users WHERE email = $1;`, u.Email)
	err = row.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.EncryptedPassword, &u.IsAdmin, &u.IsArchived, &u.CreatedAt, &u.UpdatedAt)
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

func (db *PostgresDB) UpdateUserByID(id int, u *User) (*User, error) {
	query := `
	UPDATE users
	SET email = $1, first_name = $2, last_name = $3, encrypted_password = $4, is_admin = $5, is_archived = $6, updated_at = $7`
	if u == nil {
		return nil, fmt.Errorf("user cannot be nil")
	}

	_, err := db.DB.Exec(query, u.Email, u.FirstName, u.LastName, u.EncryptedPassword, u.IsAdmin, u.IsArchived, u.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %v", err)
	}

	user, err := db.GetUserByID(id)
	if err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}

	return user, nil
}

func (db *PostgresDB) GetUserByID(id int) (*User, error) {
	user := new(User)
	query := `SELECT * FROM users WHERE id = $1;`
	row := db.DB.QueryRow(query, id)
	err := row.Scan(&user.ID, &user.Email, &user.FirstName, &user.LastName, &user.EncryptedPassword, &user.IsAdmin, &user.IsArchived, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("error scanning user: %v", err)
	}

	return user, nil
}

