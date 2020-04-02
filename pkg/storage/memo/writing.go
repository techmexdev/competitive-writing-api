package memo

import "github.com/techmexdev/competitive_writing_api/pkg/writing"

type writingStorage struct {
	data *StatefulData
}

// NewWritingStorage stores writings.
func NewWritingStorage(data *StatefulData) writing.Storage {
	return &writingStorage{data}
}

// CreateWriting adds a writing to shared stateful data: s.data.
func (s writingStorage) CreateWriting(wri writing.Writing) error {
	s.data.writings = append(s.data.writings, wri)
	return nil
}

// ReadWriting looks for writing with id in shared stateful data: s.data.
func (s writingStorage) ReadWriting(id string) (writing.Writing, error) {
	for _, w := range s.data.writings {
		if w.ID == id {
			return w, nil
		}
	}

	return writing.Writing{}, nil
}

// ListWritingsFromUser looks for writings where w.Author = usr.ID in shared stateful data: s.data.
func (s writingStorage) ListWritingsFromUser(usr writing.User) ([]writing.Writing, error) {
	return []writing.Writing{}, ErrNotImplemented
}

// ListWritingsWithPassage looks for writings where w.Passage = pass in shared stateful data: s.data.
func (s writingStorage) ListWritingsWithPassage(pass writing.Passage) ([]writing.Writing, error) {
	return []writing.Writing{}, ErrNotImplemented

}

// ListWritingsWithPassage looks for writings that have not been part of a selection in shared stateful data: s.data.
func (s writingStorage) ListUnreviewedWritingsWithPassage(pass writing.Passage) ([]writing.Writing, error) {
	return []writing.Writing{}, ErrNotImplemented

}
