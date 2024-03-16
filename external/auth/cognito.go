package auth

import (
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type cognitoAuth struct {
	session *session.Session
	client  *cognitoidentityprovider.CognitoIdentityProvider
}

func NewCognitoAuth(session *session.Session) *cognitoAuth {
	return &cognitoAuth{session: session}
}

func (c *cognitoAuth) clientCognito() {
	c.client = cognitoidentityprovider.New(c.session)
}

func (c *cognitoAuth) InitiateAuth(input *cognitoidentityprovider.InitiateAuthInput) (*cognitoidentityprovider.InitiateAuthOutput, error) {
	if c.client == nil {
		c.clientCognito()
	}

	return c.client.InitiateAuth(input)
}

func (c *cognitoAuth) ListUsers(input *cognitoidentityprovider.ListUsersInput) (*cognitoidentityprovider.ListUsersOutput, error) {
	if c.client == nil {
		c.clientCognito()
	}
	return c.client.ListUsers(input)
}

func (c *cognitoAuth) AdminGetUser(input *cognitoidentityprovider.AdminGetUserInput) (*cognitoidentityprovider.AdminGetUserOutput, error) {
	if c.client == nil {
		c.clientCognito()
	}
	return c.client.AdminGetUser(input)
}
