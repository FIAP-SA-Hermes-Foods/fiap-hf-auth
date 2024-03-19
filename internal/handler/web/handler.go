package web

import (
	"encoding/json"
	"fiap-hf-auth/internal/core/application"
	"fiap-hf-auth/internal/core/domain/entity/dto"
	"fmt"
	"net/http"

	awsEvents "github.com/aws/aws-lambda-go/events"
)

type HandlerAuth interface {
	Auth(req awsEvents.APIGatewayProxyRequest) (awsEvents.APIGatewayProxyResponse, error)
}

type handlerAuth struct {
	app application.HermesFoodsAuthApp
}

func NewHandler(app application.HermesFoodsAuthApp) *handlerAuth {
	return &handlerAuth{app: app}
}

func (h *handlerAuth) Auth(req awsEvents.APIGatewayProxyRequest) (awsEvents.APIGatewayProxyResponse, error) {

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
