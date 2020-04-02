package auth_test

import (
	"os"
	"testing"

	"github.com/techmexdev/competitive_writing_api/pkg/auth"
	"github.com/techmexdev/competitive_writing_api/pkg/auth/auth0"
)

func TestSignup(t *testing.T) {
	auth0Config := auth0.Config{
		ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
		ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
		Audience:     os.Getenv("AUTH0_AUDIENCE"),
		DB:           os.Getenv("AUTH0_DB"),
	}
	a0, err := auth0.New(auth0Config)
	if err != nil {
		t.Fatal(err)
	}
	aa := []auth.Service{a0}

	for _, a := range aa {
		_, err := a.Signup(auth.Creds{Email: "lololo@lolo.com", Pwd: "p)3icdassword123CKMDA_034"})
		if err != nil {
			t.Fatal(err)
		}
	}
}
