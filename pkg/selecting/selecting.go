package selecting

// Selection is a user's choice of best review out of reviews.
type Selection struct {
	ID             string `json:"id"`
	Selector       string `json:"selector"`
	Passage        `json:"passage"`
	BestWriting    Writing   `json:"best_analysis"`
	WritingChoices []Writing `json:"analysis_choices"`
	Review         string    `json:"review"`
}

// Storage stores Selection.
type Storage interface {
	CreateSelection(sel Selection) error
	ReadSelection(id string) (Selection, error)
	ListSelectionsFromUser(usr User) ([]Selection, error)
	ListSelectionsWithWritings(ww []Writing) ([]Selection, error)
}

// Service can manage selections.
type Service interface {
	Storage
}

type service struct {
	Storage
}

// NewService uses store as it's storage layer.
func NewService(store Storage) Service {
	return service{store}
}

// Passage is all the selection's writings passages
type Passage struct {
	ID string `json:"id"`
}

// Writing is a user's re-write/commentary on a passagraph.
type Writing struct {
	ID string `json:"id"`
}

// User is an author of a Writing.
type User struct {
	Username string `json:"username"`
}
