package postgres

import (
	"strconv"

	"github.com/jinzhu/gorm"
	"github.com/techmexdev/competitive_writing_api/pkg/passage"
)

type gormPassage struct {
	gorm.Model
	Author string
	Book   string
	Text   string
}

// PassageStorage implements writing.PassageStorage.
type PassageStorage struct {
	*gorm.DB
}

// NewPassageStorage implements writing.PassageStorage.
func NewPassageStorage(db *gorm.DB) passage.Storage {
	db.AutoMigrate(&gormPassage{})
	return &PassageStorage{db}
}

// CreatePassage adds psg to db.
func (db *PassageStorage) CreatePassage(psg passage.Passage) error {
	p := toGormPassage(psg)
	if err := db.Create(&p).Error; err != nil {
		return err
	}
	return nil

}

// ReadPassage looks for psg.
func (db *PassageStorage) ReadPassage(id string) (passage.Passage, error) {
	var psg gormPassage
	if err := db.First(&psg, id).Error; err != nil {
		return passage.Passage{}, err
	}

	return toDomainPassage(psg), nil
}

// ListPassages returns all passages.
func (db *PassageStorage) ListPassages() ([]passage.Passage, error) {
	var pp []gormPassage
	if err := db.Find(&pp).Error; err != nil {
		return []passage.Passage{}, err
	}

	return toDomainPassages(pp), nil
}

// ListPassagesFrom returns all passages where p.Author == author
func (db *PassageStorage) ListPassagesFrom(authorID string) ([]passage.Passage, error) {
	var pp []gormPassage
	db.Where(&gormPassage{Author: authorID}).Find(&pp)
	return toDomainPassages(pp), nil
}

func toGormPassage(psg passage.Passage) gormPassage {
	return gormPassage{
		Author: psg.Author,
		Book:   psg.Book,
		Text:   psg.Text,
	}
}

func toDomainPassages(psgs []gormPassage) []passage.Passage {
	var pp []passage.Passage
	for _, p := range psgs {
		pp = append(pp, toDomainPassage(p))
	}
	return pp
}

func toDomainPassage(psg gormPassage) passage.Passage {
	return passage.Passage{
		ID: strconv.Itoa(int(psg.ID)),
	}
}
