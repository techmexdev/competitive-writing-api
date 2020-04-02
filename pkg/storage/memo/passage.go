package memo

import (
	"errors"

	"github.com/techmexdev/competitive_writing_api/pkg/passage"
)

// PassageStorage implements writing.PassageStorage
type PassageStorage struct {
	data *StatefulData
}

// NewPassageStorage implements writing.PassageStorage
func NewPassageStorage(data *StatefulData) passage.Storage {
	return &PassageStorage{data: data}
}

// CreatePassage adds psg to s.passages.
func (s *PassageStorage) CreatePassage(psg passage.Passage) error {
	s.data.passages = append(s.data.passages, psg)
	return nil
}

// ReadPassage looks for psg in s.passages.
func (s *PassageStorage) ReadPassage(id string) (passage.Passage, error) {
	for _, p := range s.data.passages {
		if p.ID == id {
			return p, nil
		}
	}
	return passage.Passage{}, errors.New("passage not found")
}

// ListPassages returns all passages.
func (s *PassageStorage) ListPassages() ([]passage.Passage, error) {
	return s.data.passages, nil
}

// ListPassagesFrom returns all passages where p.Author == author
func (s *PassageStorage) ListPassagesFrom(authorID string) ([]passage.Passage, error) {
	allPP, err := s.ListPassages()
	if err != nil {
		return []passage.Passage{}, err
	}

	var pp []passage.Passage
	for _, p := range allPP {
		if p.Author == authorID {
			pp = append(pp, p)
		}
	}

	return pp, nil
}
