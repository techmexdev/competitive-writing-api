package passage

// Passage is text that users write a review of.
type Passage struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Book   string `json:"book"`
	Text   string `json:"text"`
}

// Storage stores Passages.
type Storage interface {
	CreatePassage(pass Passage) error
	ReadPassage(id string) (Passage, error)
	ListPassages() ([]Passage, error)
	ListPassagesFrom(author string) ([]Passage, error)
}

// Service can select, list and find passages.
type Service interface {
	Storage
}

type service struct {
	Storage
}

// NewService manages passages.
func NewService(store Storage) Service {
	return &service{store}
}
