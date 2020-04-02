package memo

import (
	"errors"

	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/selecting"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
)

// ErrNotImplemented is thrown during TDD.
var ErrNotImplemented = errors.New("not implemented")

// StatefulData is data that can be shared between memo storages.
type StatefulData struct {
	passages   []passage.Passage
	users      []user.User
	writings   []writing.Writing
	selections []selecting.Selection
}

// Storage implements many services' storages.
type Storage interface {
	user.Storage
	passage.Storage
	writing.Storage
	selecting.Storage
}

type usrStore user.Storage
type psgStore passage.Storage
type wriStore writing.Storage
type selStore selecting.Storage

type storage struct {
	usrStore
	psgStore
	wriStore
	selStore
}

// New creates many storages, with shared stateful data.
func New() Storage {
	d := StatefulData{users: []user.User{}}
	usr := NewUserStorage(&d)
	psg := NewPassageStorage(&d)
	wri := NewWritingStorage(&d)
	sel := NewSelectionStorage(&d)
	s := &storage{usr, psg, wri, sel}
	return s
}
