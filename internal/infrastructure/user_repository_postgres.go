package infrastructure

import (
	"database/sql"

	"github.com/vantonietti/gin-testing/internal/entity"
	"github.com/vantonietti/gin-testing/internal/repository"
)

type UserRepositoryPostgres struct {
	db *sql.DB
}

func NewUserRepositoryPostgres(db *sql.DB) repository.UserRepository {
	return &UserRepositoryPostgres{db: db}
}

func (r *UserRepositoryPostgres) Create(user *entity.User) error {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"
	return r.db.QueryRow(query, user.Name, user.Email).Scan(&user.ID)
}

func (r *UserRepositoryPostgres) GetByID(id int64) (*entity.User, error) {
	var user entity.User
	query := "SELECT id, name, email FROM users WHERE id = $1"
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepositoryPostgres) GetAll() ([]entity.User, error) {
	query := "SELECT id, name, email FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []entity.User
	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
