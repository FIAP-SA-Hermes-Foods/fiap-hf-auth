package useCase

import (
	"errors"
	"fiap-hf-auth/internal/core/domain/entity"
)

type UserUseCase interface {
	GetUserByCPF(cpf string) error
	GetUserByEmail(email string) error
	SaveUser(entity.User) error
}

type userUseCase struct {
}

func NewUserUseCase() userUseCase {
	return userUseCase{}
}

func (c userUseCase) SaveUser(user entity.User) error {
	if err := user.CPF.Validate(); err != nil {
		return err
	}

	if len(user.Email) == 0 && len(user.CPF.Value) == 0 {
		return errors.New("user's email or cpf should be not empty")
	}

	return nil

}

func (c userUseCase) GetUserByCPF(cpf string) error {
	if len(cpf) < 1 {
		return errors.New("cpf is null or invalid")
	}

	return nil
}

func (c userUseCase) GetUserByEmail(email string) error {
	if len(email) < 1 {
		return errors.New("email is null or invalid")
	}
	return nil
}
