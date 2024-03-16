package main

import (
	"fiap-hf-auth/external/auth"
	"fiap-hf-auth/external/db/dynamo"
	"fiap-hf-auth/external/jwt"
	reponosql "fiap-hf-auth/internal/adapters/driven/repositories/nosql"
	adapterAuth "fiap-hf-auth/internal/adapters/driver/auth"
	"fiap-hf-auth/internal/core/application"
	"fiap-hf-auth/internal/core/useCase"
	"fiap-hf-auth/internal/handler/web"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/marcos-dev88/genv"
)

func init() {
	if err := genv.New(); err != nil {
		log.Printf("error to define envs: %v", err)
	}
}

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
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

	app := application.NewApplication(userAuth, repo, useCase)

	jwtCall := jwt.New(
		os.Getenv("JWT_ISSUER"),
		os.Getenv("JWT_USERNAME"),
		time.Hour,
	)

	handler := web.NewHandler(app, jwtCall)
	lambda.Start(handler)
}
