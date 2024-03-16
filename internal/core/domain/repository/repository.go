package repository

import (
	"fiap-hf-auth/internal/core/domain/entity/dto"
)

type UserRepositoryNoSQL interface {
	GetUserByEmail(email string) (*dto.UserNoSQL, error)
	GetUserByCPF(cpf string) (*dto.UserNoSQL, error)
	SaveUser(client dto.UserNoSQL) (*dto.UserNoSQL, error)
}
