package repositories

import (
	"database/sql"
	"vaccine-api/models"
	"vaccine-api/utils"
)

type UserRepository struct {
	Db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) GetUserByEmail(email string) (models.User, error) {
	var user models.User

	query := `SELECT id, name, email, password FROM users WHERE email = $1`
	row := repo.Db.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (repo *UserRepository) CreateUser(user models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3)`
	_, err = repo.Db.Exec(query, user.Name, user.Email, hashedPassword)
	return err
}
