package memo

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
)

// userStorage implements user.Storage
type userStorage struct {
	data *StatefulData
}

// NewUserStorage implements user.Storage
func NewUserStorage(data *StatefulData) user.Storage {
	return &userStorage{data}
}

// CreateUser adds user to s.users
func (s *userStorage) CreateUser(authID string) (user.User, error) {
	u := user.User{ID: uuid.NewV4().String(), AuthID: authID}
	s.data.users = append(s.data.users, u)
	return u, nil
}

// ReadUser looks for User with id in s.users
func (s *userStorage) ReadUser(authID string) (user.User, error) {
	for _, u := range s.data.users {
		if u.AuthID == authID {
			return u, nil
		}
	}
	return user.User{}, errors.New("user not found")
}

func (s *userStorage) ReadWritingAuthor(wri user.Writing) ([]user.User, error) {
	return []user.User{}, ErrNotImplemented
}

func (s *userStorage) ReadSelectionAuthor(sel user.Selection) ([]user.User, error) {
	return []user.User{}, ErrNotImplemented
}
