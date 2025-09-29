package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// A chave secreta usada para assinar os tokens.
// Em um ambiente de produção, isso deveria ser uma variável de ambiente.
var jwtKey = []byte("miguelma3")

// Claims define a estrutura dos dados que serão armazenados no token.
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Faz o parse do token, verificando a assinatura e a validade.
	// A função anônima fornece a chave para a verificação da assinatura.
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token invalido")
	}

	return claims, nil
}
