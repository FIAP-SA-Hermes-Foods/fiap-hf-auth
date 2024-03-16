package web

import (
	"encoding/json"
	"fiap-hf-auth/external/jwt"
	"fiap-hf-auth/internal/core/application"
	"fiap-hf-auth/internal/core/domain/entity/dto"
	"fmt"
	"net/http"
	"os"

	awsEvents "github.com/aws/aws-lambda-go/events"
)

type HandlerAuth interface {
	Auth(req awsEvents.APIGatewayProxyRequest) (awsEvents.APIGatewayProxyResponse, error)
}

type handlerAuth struct {
	app   application.HermesFoodsAuthApp
	jwtHf jwt.JwtHF
}

func NewHandler(app application.HermesFoodsAuthApp, jwtHf jwt.JwtHF) *handlerAuth {
	return &handlerAuth{app: app, jwtHf: jwtHf}
}

func (h *handlerAuth) Auth(req awsEvents.APIGatewayProxyRequest) (awsEvents.APIGatewayProxyResponse, error) {

	if _, ok := req.Headers["Auth-token"]; !ok {
		return awsEvents.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body:            `{"error": "auth-token is invalid or null"}`,
			IsBase64Encoded: false,
		}, nil
	}

	statusCode, err := h.jwtHf.ValidateToken(req.Headers["Auth-token"], []byte(os.Getenv("JWT_SIGIN_KEY")))

	if err != nil {
		return awsEvents.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body:            fmt.Sprintf(`{"error": "%s"}`, err.Error()),
			IsBase64Encoded: false,
		}, err
	}

	if statusCode != 200 {
		return awsEvents.APIGatewayProxyResponse{
			StatusCode: http.StatusUnauthorized,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body:            `{"error": "auth-token is invalid or null"}`,
			IsBase64Encoded: false,
		}, nil
	}

	var in dto.Input

	if err := json.Unmarshal([]byte(req.Body), &in); err != nil {
		return awsEvents.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body:            fmt.Sprintf(`{"error": "%s"}`, err.Error()),
			IsBase64Encoded: false,
		}, err
	}

	out, err := h.app.AuthUser(in)

	if err != nil {
		return awsEvents.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body: fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, err
	}

	outJson, err := json.Marshal(out)

	if err != nil {
		return awsEvents.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Headers: map[string]string{
				"content-type": "application/json",
			},
			Body: fmt.Sprintf(`{"error": "%s"}`, err.Error()),
		}, err

	}

	return awsEvents.APIGatewayProxyResponse{
		StatusCode: out.StatusCode,
		Headers: map[string]string{
			"content-type": "application/json",
		},
		Body: string(outJson),
	}, nil

}
