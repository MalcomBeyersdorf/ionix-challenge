package repositories

import (
	"regexp"
	"testing"
	"vaccine-api/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestGetUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "password"}).
		AddRow(1, "John Doe", "john@example.com", "hashedPassword")

	mock.ExpectQuery(regexp.QuoteMeta("SELECT id, name, email, password FROM users WHERE email = $1")).
		WithArgs("john@example.com").
		WillReturnRows(rows)

	user, err := repo.GetUserByEmail("john@example.com")
	if err != nil {
		t.Errorf("error was not expected while getting user: %s", err)
	}

	if user.Email != "john@example.com" {
		t.Errorf("expected user with email john@example.com, got %v", user.Email)
	}
}

func TestCreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewUserRepository(db)

	mock.ExpectExec(regexp.QuoteMeta("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)")).
		WithArgs("Jane Doe", "jane@example.com", sqlmock.AnyArg()). // Usar AnyArg() para simular la contraseña hasheada
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.CreateUser(models.User{Name: "Jane Doe", Email: "jane@example.com", Password: "password123"})
	if err != nil {
		t.Errorf("error was not expected while creating user: %s", err)
	}
}
