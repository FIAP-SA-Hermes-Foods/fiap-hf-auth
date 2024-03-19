package application

import (
	"encoding/json"
	"fiap-hf-auth/external/auth"
	"fiap-hf-auth/external/db/dynamo"
	reponosql "fiap-hf-auth/internal/adapters/driven/repositories/nosql"
	adapterAuth "fiap-hf-auth/internal/adapters/driver/auth"
	"log"
	"os"

	"fiap-hf-auth/internal/core/domain/entity/dto"
	"fiap-hf-auth/internal/core/useCase"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

// go test -v -count=1 -failfast -run ^Test_App$
func Test_App(t *testing.T) {
	configAws := aws.NewConfig()
	configAws.Region = aws.String("us-east-1")

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config:            *configAws,
		SharedConfigState: session.SharedConfigEnable,
	}))

	db := dynamo.NewDynamoDB(sess)

	cognito := auth.NewCognitoAuth(sess)

	userAuth := adapterAuth.NewUserAuth(
		os.Getenv("USER_POOL_ID"),
		os.Getenv("USER_POOL_CLIENT_ID"),
		cognito,
	)

	repo := reponosql.NewUserDynamoDB(db, os.Getenv("DB_TABLE"))

	useCase := useCase.NewUserUseCase()

	app := NewApplication(userAuth, repo, useCase)

	in := dto.Input{
		User: &dto.UserInput{
			CPF:          "some",
			Email:        "",
			Password:     "test",
			WantRegister: true,
		},
	}

	out, err := app.AuthUser(in)

	if err != nil {
		log.Println(err)
		return
	}

	o, _ := json.Marshal(out)

	log.Println(string(o))

}
