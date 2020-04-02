package postgres

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/techmexdev/competitive_writing_api/pkg/passage"
	"github.com/techmexdev/competitive_writing_api/pkg/selecting"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
)

// ErrNotImplemented is thrown during TDD.
var ErrNotImplemented = errors.New("not implemented")

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
func New(db *gorm.DB) Storage {
	usr := NewUserStorage(db)
	psg := NewPassageStorage(db)
	wri := NewWritingStorage(db)
	sel := NewSelectingStorage(db)
	s := &storage{usr, psg, wri, sel}
	return s
}
