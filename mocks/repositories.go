package mocks

import (
	"fiap-hf-auth/internal/core/domain/entity/dto"
	"strings"
)

type MockUserRepositoryNoSQL struct {
	WantOut *dto.UserNoSQL
	WantErr error
}

func (m MockUserRepositoryNoSQL) GetUserByEmail(email string) (*dto.UserNoSQL, error) {
	if m.WantErr != nil && strings.EqualFold("err GetUserByEmail", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantOut, nil
}

func (m MockUserRepositoryNoSQL) GetUserByCPF(cpf string) (*dto.UserNoSQL, error) {
	if m.WantErr != nil && strings.EqualFold("errGetUserByCPF", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantOut, nil
}

func (m MockUserRepositoryNoSQL) SaveUser(user dto.UserNoSQL) (*dto.UserNoSQL, error) {
	if m.WantErr != nil && strings.EqualFold("errSaveUser", m.WantErr.Error()) {
		return nil, m.WantErr
	}

	return m.WantOut, nil
}
