package jwt

import (
	"fmt"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/techmexdev/competitive_writing_api/pkg/auth"
)

type service struct {
	signature string
}

// New creates a firebase client.
func New(signature string) auth.Service {
	return &service{signature: signature}
}

func (s service) Signup(creds auth.Creds) (token string, err error) {
	// Create the Claims
	claims := &jwt.StandardClaims{
		Issuer:  "Test",
		Subject: creds.Email,
	}

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := tok.SignedString([]byte(s.signature))
	if err != nil {
		return "", fmt.Errorf("failed signing key: %s", err)
	}

	return ss, nil
}

func (s service) Login(creds auth.Creds) (token string, err error) {
	panic("not implemented") // TODO: Implement
}

// Verify is a replacement for firebaseAuth.VerifyIDToken.
func (s service) Verify(tokenStr string) (username string, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(s.signature), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		log.Printf("claims = %+v\n", claims)
		return claims["subject"].(string), nil
	}

	return "", fmt.Errorf("bad token: %s", err)
}
