package db

import (
    "testing"

    "github.com/aaron-smits/templ-starter/model"
    "github.com/DATA-DOG/go-sqlmock"
    "github.com/stretchr/testify/assert"
)

func TestAddTodo(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("INSERT INTO todos").
        WithArgs("Test title", "Test body", false, "Test user").
        WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

    pgDB := &PostgresDB{DB: db}

    err = pgDB.AddTodo(model.Todo{
        Title:  "Test title",
        Body:   "Test body",
        Done:   false,
        UserID: "Test user",
    })

    assert.NoError(t, err)
    assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDeleteTodo(t *testing.T) {
    db, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer db.Close()

    mock.ExpectQuery("DELETE FROM todos").
		WithArgs("1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

    pgDB := &PostgresDB{DB: db}

    err = pgDB.DeleteTodo("1")

    assert.NoError(t, err)
    assert.NoError(t, mock.ExpectationsWereMet())
}
