package dto

type UserInput struct {
	CPF          string `json:"cpf,omitempty"`
	Email        string `json:"email,omitempty"`
	Password     string `json:"password,omitempty"`
	WantRegister bool   `json:"wantRegister"`
}

type (
	UserNoSQL struct {
		UUID      string `json:"uuid"`
		Username  string `json:"username"`
		Password  string `json:"password"`
		Name      string `json:"name,omitempty"`
		CPF       string `json:"cpf,omitempty"`
		Email     string `json:"email,omitempty"`
		CreatedAt string `json:"created_at,omitempty"`
	}
)

type (
	UserPool struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Name     string `json:"name,omitempty"`
		CPF      string `json:"cpf,omitempty"`
		Email    string `json:"email,omitempty"`
	}

	UserPoolOut struct {
		Name          string `json:"name"`
		Username      string `json:"username"`
		CPF           string `json:"cpf"`
		Email         string `json:"email"`
		EmailVerified string `json:"email_verified"`
	}
)
