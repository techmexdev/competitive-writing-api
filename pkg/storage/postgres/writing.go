package postgres

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/techmexdev/competitive_writing_api/pkg/writing"
)

type writingStorage struct {
	*gorm.DB
}

type gormWriting struct {
	gorm.Model
	Author  string
	Passage gormPassage
	Text    string
}

// NewWritingStorage stores writings.
func NewWritingStorage(db *gorm.DB) writing.Storage {
	db.AutoMigrate(&gormWriting{})
	return &writingStorage{db}
}

func (db *writingStorage) CreateWriting(wri writing.Writing) (writing.Writing, error) {
	w := toGormWriting(wri)
	var written gormWriting
	err := db.Create(&w).First(&written).Error
	if err != nil {
		return writing.Writing{}, err
	}

	return toDomainWriting(written), nil
}

func (db *writingStorage) ReadWriting(id string) (writing.Writing, error) {
	var w gormWriting
	err := db.First(&w, id).Error
	if err != nil {
		return writing.Writing{}, err
	}

	return toDomainWriting(w), nil
}

func (db *writingStorage) ListWritingsFromUser(usr writing.User) ([]writing.Writing, error) {
	var ww []gormWriting
	err := db.Where(&gormWriting{Author: usr.ID}).Find(&ww).Error
	if err != nil {
		return []writing.Writing{}, err
	}

	return toDomainWritings(ww), nil
}

func (db *writingStorage) ListWritingsWithPassage(psg writing.Passage) ([]writing.Writing, error) {
	var ww []gormWriting
	err := db.Where("passage_id = ?", psg.ID).Find(&ww).Error
	if err != nil {
		return []writing.Writing{}, err
	}

	return toDomainWritings(ww), nil
}

func (db *writingStorage) ListUnreviewedWritingsWithPassage(pass writing.Passage) ([]writing.Writing, error) {
	panic("not implemented") // TODO: Implement
}

func toGormWriting(w writing.Writing) gormWriting {
	id, _ := strconv.Atoi(w.Passage.ID)
	return gormWriting{
		Author:  w.Author,
		Passage: gormPassage{Model: gorm.Model{ID: uint(id)}},
		Text:    w.Text,
	}
}

func toDomainWritings(ww []gormWriting) []writing.Writing {
	var wris []writing.Writing
	for _, w := range ww {
		wris = append(wris, toDomainWriting(w))
	}

	return wris
}

func toDomainWriting(w gormWriting) writing.Writing {
	return writing.Writing{
		ID:     strconv.Itoa(int(w.ID)),
		Author: w.Author,
		Passage: writing.Passage{
			ID: strconv.Itoa(int(w.Passage.ID)),
		},
		Text: w.Text,
	}
}
