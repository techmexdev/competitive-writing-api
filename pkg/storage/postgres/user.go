package postgres

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/techmexdev/competitive_writing_api/pkg/user"
)

type userStorage struct {
	*gorm.DB
}

type gormUser struct {
	gorm.Model
	AuthID string `json:"auth_id"`
	Score  `json:"score"`
}

// Score shows how well a user's writings have been.
type Score struct {
	TotalPoints int `json:"total_points"`
}

// NewUserStorage stores users.
func NewUserStorage(db *gorm.DB) user.Storage {
	db.AutoMigrate(&gormUser{})
	return &userStorage{db}
}

func (db *userStorage) CreateUser(authID string) (user.User, error) {
	u := gormUser{AuthID: authID}

	err := db.Create(&u).Error
	if err != nil {
		return user.User{}, err
	}

	usr, err := db.ReadUser(authID)
	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

func (db *userStorage) ReadUser(authID string) (user.User, error) {
	usr := gormUser{}
	err := db.Where(&gormUser{AuthID: authID}).First(&usr).Error
	if err != nil {
		return user.User{}, err
	}

	return toDomainUser(usr), nil
}

func (db *userStorage) ReadWritingAuthor(wri user.Writing) ([]user.User, error) {
	panic("not implemented") // TODO: Implement
}

func (db *userStorage) ReadSelectionAuthor(sel user.Selection) ([]user.User, error) {
	panic("not implemented") // TODO: Implement
}

func toGormUser(usr user.User) gormUser {
	return gormUser{
		AuthID: usr.AuthID,
		Score: Score{
			TotalPoints: usr.TotalPoints,
		},
	}
}

func toDomainUser(usr gormUser) user.User {
	return user.User{
		ID:     strconv.Itoa(int(usr.ID)),
		AuthID: usr.AuthID,
		Score: user.Score{
			TotalPoints: usr.TotalPoints,
		},
	}
}
