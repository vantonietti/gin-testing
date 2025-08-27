package repository

import "github.com/vantonietti/gin-testing/internal/entity"

type UserRepository interface {
	Create(user *entity.User) error
	GetByID(id int64) (*entity.User, error)
	GetAll() ([]entity.User, error)
}
