package application

import (
	"errors"
	"fiap-hf-auth/internal/core/domain/auth"
	"fiap-hf-auth/internal/core/domain/entity"
	"fiap-hf-auth/internal/core/domain/entity/dto"
	"fiap-hf-auth/internal/core/domain/repository"
	valueobject "fiap-hf-auth/internal/core/domain/valueObject"
	"fiap-hf-auth/internal/core/useCase"
	"time"

	"github.com/google/uuid"
)

type HermesFoodsAuthApp interface {
	AuthUser(in dto.Input) (*dto.Output, error)
}

type hermesFoodsAuthApp struct {
	userAuth      auth.UserAuth
	userNoSQLRepo repository.UserRepositoryNoSQL
	userUseCase   useCase.UserUseCase
}

func NewApplication(userAuth auth.UserAuth, userRepo repository.UserRepositoryNoSQL, userUseCase useCase.UserUseCase) hermesFoodsAuthApp {
	return hermesFoodsAuthApp{userAuth: userAuth, userNoSQLRepo: userRepo, userUseCase: userUseCase}

}

func (h hermesFoodsAuthApp) AuthUser(in dto.Input) (*dto.Output, error) {
	if !in.User.WantRegister {
		return &dto.Output{StatusCode: 200, Message: "OK"}, nil
	}

	if in.User == nil {
		return &dto.Output{StatusCode: 401, Message: "username or password is incorrect"}, nil
	}

	if len(in.User.CPF) == 0 && len(in.User.Email) == 0 {
		return &dto.Output{StatusCode: 401, Message: "username or password is incorrect"}, nil
	}

	var userFound *dto.UserPool

	if len(in.User.Email) == 0 {
		found, err := h.authUserByCPF(in)

		if err != nil {
			return &dto.Output{StatusCode: 500, Message: err.Error()}, nil
		}
		userFound = found
	} else {
		found, err := h.authUserByEmail(in)

		if err != nil {
			return &dto.Output{StatusCode: 500, Message: err.Error()}, nil
		}
		userFound = found
	}

	if userFound == nil {
		return &dto.Output{StatusCode: 404, Message: "user not found, try signup through this link: URL HERE"}, nil
	}

	if userFound.Password != in.User.Password {
		return &dto.Output{StatusCode: 401, Message: "username or password is incorrect"}, nil
	}

	return &dto.Output{StatusCode: 200, Message: "OK"}, nil
}

func (h hermesFoodsAuthApp) authUserByCPF(in dto.Input) (*dto.UserPool, error) {
	userCPFNoSQL, err := h.getUserByCPF(in)

	if err != nil {
		return nil, err
	}

	if userCPFNoSQL != nil {
		out := &dto.UserPool{
			Username: userCPFNoSQL.Username,
			Name:     userCPFNoSQL.Name,
			CPF:      userCPFNoSQL.CPF,
			Email:    userCPFNoSQL.Email,
			Password: userCPFNoSQL.Password,
		}

		return out, nil
	}

	userList, err := h.userAuth.ListUsers()

	if err != nil {
		return nil, err
	}

	var foundUserPool *dto.UserPoolOut

	for i := 0; i < len(userList); i++ {

		if userList[i].CPF == in.User.CPF {
			foundUserPool = &userList[i]
			break
		}
	}

	if foundUserPool != nil {
		if err := h.userAuth.UserAuthentication(foundUserPool.Email, in.User.Password); err != nil {
			return nil, err
		}

		saveDbUser := dto.UserPool{
			Username: foundUserPool.Username,
			Name:     foundUserPool.Name,
			CPF:      foundUserPool.CPF,
			Email:    foundUserPool.Email,
			Password: in.User.Password,
		}

		_, err := h.saveUser(saveDbUser)

		if err != nil {
			return nil, err
		}

		out := dto.UserPool{
			Username: foundUserPool.Username,
			Password: in.User.Password,
			Name:     foundUserPool.Name,
			CPF:      foundUserPool.CPF,
			Email:    foundUserPool.Email,
		}

		return &out, nil

	}

	return nil, nil
}

func (h hermesFoodsAuthApp) authUserByEmail(in dto.Input) (*dto.UserPool, error) {
	userEmailNoSQL, err := h.getUserByEmail(in)

	if err != nil {
		return nil, err
	}

	if userEmailNoSQL != nil {

		out := &dto.UserPool{
			Username: userEmailNoSQL.Username,
			Name:     userEmailNoSQL.Name,
			CPF:      userEmailNoSQL.CPF,
			Email:    userEmailNoSQL.Email,
			Password: userEmailNoSQL.Password,
		}

		return out, nil
	}

	userList, err := h.userAuth.ListUsers()

	if err != nil {
		return nil, err
	}

	var foundUserPool *dto.UserPoolOut

	for i := 0; i < len(userList); i++ {

		if userList[i].Email == in.User.Email {
			foundUserPool = &userList[i]
			break
		}
	}

	if foundUserPool != nil {
		if err := h.userAuth.UserAuthentication(in.User.Email, in.User.Password); err != nil {
			return nil, err
		}

		saveDbUser := dto.UserPool{
			Username: foundUserPool.Username,
			Name:     foundUserPool.Name,
			CPF:      foundUserPool.CPF,
			Email:    foundUserPool.Email,
			Password: in.User.Password,
		}

		_, err := h.saveUser(saveDbUser)

		if err != nil {
			return nil, err
		}

		out := dto.UserPool{
			Username: foundUserPool.Username,
			Password: in.User.Password,
			Name:     foundUserPool.Name,
			CPF:      foundUserPool.CPF,
			Email:    foundUserPool.Email,
		}

		return &out, nil
	}

	return nil, nil
}

func (h hermesFoodsAuthApp) getUserByCPF(in dto.Input) (*dto.UserPool, error) {
	if in.User == nil {
		return nil, errors.New("userInput input is null or invalid")
	}

	if err := h.userUseCase.GetUserByCPF(in.User.CPF); err != nil {
		return nil, err
	}

	noSqlOut, err := h.userNoSQLRepo.GetUserByCPF(in.User.CPF)

	if err != nil {
		return nil, err
	}

	if noSqlOut != nil {
		out := dto.UserPool{
			Name:     noSqlOut.Name,
			Username: noSqlOut.Username,
			Password: noSqlOut.Password,
			CPF:      noSqlOut.CPF,
			Email:    noSqlOut.Email,
		}

		return &out, nil
	}

	return nil, nil

}

func (h hermesFoodsAuthApp) getUserByEmail(in dto.Input) (*dto.UserPool, error) {
	if in.User == nil {
		return nil, errors.New("userInput input is null or invalid")
	}

	if err := h.userUseCase.GetUserByEmail(in.User.Email); err != nil {
		return nil, err
	}

	noSqlOut, err := h.userNoSQLRepo.GetUserByEmail(in.User.Email)

	if err != nil {
		return nil, err
	}

	if noSqlOut != nil {

		out := dto.UserPool{
			Name:     noSqlOut.Name,
			Username: noSqlOut.Username,
			Password: noSqlOut.Password,
			CPF:      noSqlOut.CPF,
			Email:    noSqlOut.Email,
		}

		return &out, nil
	}

	return nil, nil
}

func (h hermesFoodsAuthApp) saveUser(in dto.UserPool) (*dto.UserPool, error) {

	user := entity.User{
		CPF: valueobject.Cpf{
			Value: in.CPF,
		},
		Email: in.Email,
		CreatedAt: valueobject.CreatedAt{
			Value: time.Now(),
		},
	}

	if err := h.userUseCase.SaveUser(user); err != nil {
		return nil, err
	}

	cNoSQL := dto.UserNoSQL{
		UUID:      uuid.NewString(),
		Name:      in.Name,
		Username:  in.Username,
		Password:  in.Password,
		CPF:       in.CPF,
		Email:     in.Email,
		CreatedAt: user.CreatedAt.Format(),
	}

	outNoSql, err := h.userNoSQLRepo.SaveUser(cNoSQL)

	if err != nil {
		return nil, err
	}

	out := dto.UserPool{
		Username: outNoSql.Username,
		Password: outNoSql.Password,
		Name:     outNoSql.Name,
		CPF:      outNoSql.CPF,
		Email:    outNoSql.Email,
	}

	return &out, nil
}
