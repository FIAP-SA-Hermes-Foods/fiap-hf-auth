package auth

import (
	"fiap-hf-auth/internal/core/domain/entity/dto"
)

type UserAuth interface {
	ListUsers() ([]dto.UserPoolOut, error)
	UserAuthentication(username, password string) error
}
