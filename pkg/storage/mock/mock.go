package mock

import cw "github.com/techmexdev/competitive_writing_api"

// New creates Storage with hard-coded data stored.
func New() cw.Storage {
	return cw.Storage{
		User:      &userStorage{users: uu},
		Passage:   &passageStorage{passages: pp},
		Analysis:  &analysisStorage{analyses: aa},
		Selection: &selectionStorage{selections: ss},
	}
}

type userStorage struct {
	users []cw.User
}

func (s *userStorage) Create(usr cw.User) (err error) {
	return nil
}

func (s *userStorage) Read(username string) (usr cw.User, err error) {
	return s.users[0], nil
}

func (s *userStorage) List() (usr []cw.User, err error) {
	return s.users, nil
}

type passageStorage struct {
	passages []cw.Passage
}

func (s *passageStorage) Create(par cw.Passage) (err error) {
	return nil
}

func (s *passageStorage) Read(id string) (par cw.Passage, err error) {
	return s.passages[0], nil
}

func (s *passageStorage) List() (par []cw.Passage, err error) {
	return s.passages, nil
}

type analysisStorage struct {
	analyses []cw.Analysis
}

func (s *analysisStorage) Create(par cw.Analysis) (err error) {
	return nil
}

func (s *analysisStorage) Read(id string) (par cw.Analysis, err error) {
	return s.analyses[0], nil
}

func (s *analysisStorage) List() (par []cw.Analysis, err error) {
	return s.analyses, nil
}

type selectionStorage struct {
	selections []cw.Selection
}

func (s *selectionStorage) Create(par cw.Selection) (err error) {
	return nil
}

func (s *selectionStorage) Read(id string) (par cw.Selection, err error) {
	return s.selections[0], nil
}

func (s *selectionStorage) List() (par []cw.Selection, err error) {
	return s.selections, nil
}
