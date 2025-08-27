package usecase

import (
	"github.com/vantonietti/gin-testing/internal/entity"
	"github.com/vantonietti/gin-testing/internal/repository"
)

type UserUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (uc *UserUsecase) CreateUser(user *entity.User) error {
	if user.Email == "" {
		return ErrInvalidEmail
	}
	return uc.repo.Create(user)
}

func (uc *UserUsecase) GetUserByID(id int64) (*entity.User, error) {
	return uc.repo.GetByID(id)
}

func (uc *UserUsecase) ListUsers() ([]entity.User, error) {
	return uc.repo.GetAll()
}
