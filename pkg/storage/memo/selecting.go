package memo

import "github.com/techmexdev/competitive_writing_api/pkg/selecting"

type selectionStorage struct {
	data *StatefulData
}

// NewSelectionStorage uses shared data.
func NewSelectionStorage(data *StatefulData) selecting.Storage {
	return selectionStorage{data}
}

// CreateWriting adds a writing to shared stateful data: s.data.
func (s selectionStorage) CreateSelection(sel selecting.Selection) error {
	panic("not implemented") // TODO: Implement
}

// ReadSelection looks for selections with id in shared stateful data: s.data.
func (s selectionStorage) ReadSelection(id string) (selecting.Selection, error) {
	panic("not implemented") // TODO: Implement
}

// ListSelectionsFromUser looks for selections
func (s selectionStorage) ListSelectionsFromUser(usr selecting.User) ([]selecting.Selection, error) {
	panic("not implemented") // TODO: Implement
}

// ListSelectionsWithWritings looks for selections
func (s selectionStorage) ListSelectionsWithWritings(ww []selecting.Writing) ([]selecting.Selection, error) {
	panic("not implemented") // TODO: Implement
}
