package auth

import (
	"fiap-hf-auth/internal/core/auth"
	ua "fiap-hf-auth/internal/core/domain/auth"
	"fiap-hf-auth/internal/core/domain/entity/dto"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var _ ua.UserAuth = (*userAuth)(nil)

type userAuth struct {
	userPoolID string
	clientID   string
	authUsers  auth.Auth
}

func NewUserAuth(userPoolID, clientID string, auth auth.Auth) *userAuth {
	return &userAuth{userPoolID: userPoolID, clientID: clientID, authUsers: auth}
}

func (u *userAuth) ListUsers() ([]dto.UserPoolOut, error) {
	params := &cognitoidentityprovider.ListUsersInput{
		UserPoolId: aws.String(u.userPoolID),
	}

	result, err := u.authUsers.ListUsers(params)

	if err != nil {
		return nil, err
	}

	var out = make([]dto.UserPoolOut, 0)
	for _, user := range result.Users {
		var u dto.UserPoolOut

		if user.Username != nil {
			u.Username = *user.Username
		}

		for _, attribute := range user.Attributes {

			if attribute.Name != nil {

				if *attribute.Name == "name" && attribute.Value != nil {
					u.Name = *attribute.Value
				}

				if *attribute.Name == "email" && attribute.Value != nil {
					u.Email = *attribute.Value
				}

				if *attribute.Name == "email_verified" && attribute.Value != nil {
					u.EmailVerified = *attribute.Value

				}

				if *attribute.Name == "custom:cpf_cnpj" && attribute.Value != nil {
					u.CPF = *attribute.Value
				}
			}
		}

		out = append(out, u)
	}

	return out, nil
}

func (u *userAuth) UserAuthentication(username, password string) error {

	authParams := map[string]*string{
		"USERNAME": &username,
		"PASSWORD": &password,
	}

	params := &cognitoidentityprovider.InitiateAuthInput{
		AuthFlow:       aws.String("USER_PASSWORD_AUTH"),
		AuthParameters: authParams,
		ClientId:       aws.String(u.clientID),
	}

	_, err := u.authUsers.InitiateAuth(params)
	if err != nil {
		return err
	}
	return nil
}
