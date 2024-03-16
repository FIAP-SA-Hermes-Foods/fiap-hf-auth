package entity

import (
	vo "fiap-hf-auth/internal/core/domain/valueObject"
)

type User struct {
	ID        int          `json:"id,omitempty"`
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	Name      string       `json:"name,omitempty"`
	CPF       vo.Cpf       `json:"cpf,omitempty"`
	Email     string       `json:"email,omitempty"`
	CreatedAt vo.CreatedAt `json:"createdAt,omitempty"`
}
