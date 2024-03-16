package dto

type Output struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Token      string `json:"token,omitempty"`
}
