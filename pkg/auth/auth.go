package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Auth implements cw.Auth with JWTs
type Auth struct {
	Issuer           string
	HMACSampleSecret string
}

// Verify errors if token != a.GoodToken
func (a Auth) Verify(token string) (username string, err error) {
	// Parse takes the token string and a function for looking up the key. The latter is especially
	// useful if you use multiple keys for your application.  The standard is to use 'kid' in the
	// head of the token to identify which key to use, but the parsed token (head and claims) is provided
	// to the callback, providing flexibility.
	tok, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return a.HMACSampleSecret, nil
	})

	var claims jwt.MapClaims
	var ok bool
	if claims, ok = tok.Claims.(jwt.MapClaims); !ok || !tok.Valid {
		return "", fmt.Errorf("bad token: %v", tok)
	}

	var usn string
	if usn, ok = claims["usn"].(string); !ok || usn == "" {
		return "", fmt.Errorf("did not find username in token claim 'usn': %v", tok)
	}

	time.Now()
	return usn, nil
}

// CreateToken creates a jwt token signed with a.HmacSampleSecret
func (a Auth) CreateToken(username string) (token string, err error) {
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Audience:  a.Issuer,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
		IssuedAt:  time.Now().Unix(),
		Issuer:    a.Issuer,
		Subject:   username,
	})

	tokenString, err := tok.SignedString(a.HMACSampleSecret)
	return tokenString, nil
}
