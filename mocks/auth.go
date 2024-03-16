package mocks

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"strings"
)

type AuthMock struct {
	WantOutListUsers    *cognitoidentityprovider.ListUsersOutput
	WantOutAdminGetUser *cognitoidentityprovider.AdminGetUserOutput
	WantOutInitiateAuth *cognitoidentityprovider.InitiateAuthOutput
	WantErr             error
}

func (a AuthMock) ListUsers(input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
	if a.WantErr != nil && strings.EqualFold("errListUsers", a.WantErr.Error()) {
		return nil, a.WantErr
	}

	return a.WantOutListUsers, nil

}

func (a AuthMock) AdminGetUser(input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	if a.WantErr != nil && strings.EqualFold("errAdminGetUser", a.WantErr.Error()) {
		return nil, a.WantErr
	}

	return a.WantOutAdminGetUser, nil
}

func (a AuthMock) InitiateAuth(input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	if a.WantErr != nil && strings.EqualFold("errInitiateAuth", a.WantErr.Error()) {
		return nil, a.WantErr
	}

	return a.WantOutInitiateAuth, nil
}
