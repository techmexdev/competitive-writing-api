package memo_test

import (
	"testing"

	"github.com/techmexdev/competitive_writing_api/pkg/storage/memo"
)

func TestUserCreate(t *testing.T) {
	store := memo.NewUserStorage(&memo.StatefulData{})
	authID := "01100101"
	u, err := store.CreateUser(authID)
	if err != nil {
		t.Fatalf("failed creating user: %s", err)
	}
	if u.AuthID != authID {
		t.Fatalf("have authID %s, want %s", u.AuthID, authID)
	}
}

func TestUserRead(t *testing.T) {
	store := memo.NewUserStorage(&memo.StatefulData{})
	authID := "01100101"
	_, err := store.CreateUser(authID)
	if err != nil {
		t.Fatalf("failed creating user: %s", err)
	}

	u, err := store.ReadUser(authID)
	if err != nil {
		t.Fatalf("failed reading user: %s", err)
	}

	if u.AuthID != authID {
		t.Fatalf("have authID %s, want %s", u.AuthID, authID)
	}
}
