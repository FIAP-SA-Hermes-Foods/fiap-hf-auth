package jwt

import (
	"net/http"
	"time"

	jwtGo "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwtGo.StandardClaims
}

type JwtHF interface {
	Encode(siginKey []byte) (string, error)
	ValidateToken(token string, siginKey []byte) (int, error)
}

type jwtHF struct {
	issuer    string
	username  string
	expiresAt time.Duration
}

func New(issuer, username string, expiresAt time.Duration) JwtHF {
	return &jwtHF{issuer: issuer, username: username, expiresAt: expiresAt}
}

func (j *jwtHF) Encode(siginKey []byte) (string, error) {
	options := Claims{
		Username: j.username,
		StandardClaims: jwtGo.StandardClaims{
			ExpiresAt: time.Now().Add(j.expiresAt).Unix(),
			Issuer:    j.issuer,
		},
	}

	token := jwtGo.NewWithClaims(jwtGo.SigningMethodHS256, options)

	strToken, err := token.SignedString(siginKey)

	if err != nil {
		return "", err
	}

	return strToken, nil
}

func (j *jwtHF) ValidateToken(tokenStr string, siginKey []byte) (int, error) {

	if len(tokenStr) == 0 {
		return http.StatusUnauthorized, nil
	}

	t, err := jwtGo.Parse(tokenStr, func(token *jwtGo.Token) (interface{}, error) {
		return siginKey, nil
	})

	if err != nil {
		return http.StatusInternalServerError, err
	}

	if t.Valid {
		return http.StatusOK, nil
	}

	return http.StatusUnauthorized, nil
}
