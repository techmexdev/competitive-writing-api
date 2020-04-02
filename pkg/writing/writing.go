package writing

// Writing is a user's re-write/commentary on a passagraph.
type Writing struct {
	ID      string `json:"id"`
	Author  string `json:"author"` // user_id
	Passage `json:"passage"`
	Text    string `json:"text"`
}

// User is an author of a Writing.
type User struct {
	ID string `json:"id"`
}

// Passage is a Writing's associated passage.
type Passage struct {
	ID string `json:"id"`
}

// Storage stores Writing.
type Storage interface {
	CreateWriting(wri Writing) (Writing, error)
	ReadWriting(id string) (Writing, error)
	ListWritingsFromUser(usr User) ([]Writing, error)
	ListWritingsWithPassage(pass Passage) ([]Writing, error)
	ListUnreviewedWritingsWithPassage(pass Passage) ([]Writing, error)
}

// Service manages writings.
type Service interface {
	Storage
	SelectNextWritingsForReview(usr User, pass Passage) (ww []Writing, err error)
}

// Service can select, list and find passages.
type service struct {
	Storage
}

// NewService uses store as it's storage layer.
func NewService(store Storage) Service {
	return &service{store}
}

func (s service) SelectNextWritingsForReview(usr User, pass Passage) ([]Writing, error) {
	var ww []Writing
	unRev, err := s.ListUnreviewedWritingsWithPassage(pass)
	if err != nil {
		return []Writing{}, err
	}

	for _, w := range unRev {
		if w.Author != usr.ID {
			ww = append(ww, w)
		}
	}

	return ww, nil
}
