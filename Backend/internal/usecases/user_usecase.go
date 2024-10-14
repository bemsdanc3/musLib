package usecases

import (
	"Backend/internal/entities"
	"Backend/internal/repository"
)

type UserUsecase interface {
	GetAllUsers() ([]*entities.User, error)
	CreateUSer(user *entities.User) error
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUseCase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) GetAllUsers() ([]*entities.User, error) {
	return u.repo.GetAllUsers()
}

func (u *userUsecase) CreateUSer(user *entities.User) error {
	return u.repo.CreateUSer(user)
}
