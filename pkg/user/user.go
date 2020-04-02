package user

// User is a user of the app.
type User struct {
	ID     string `json:"id"`
	AuthID string `json:"auth_id"`
	Score  `json:"score"`
}

// Writing is authored by a user.
type Writing struct {
	ID string `json:"id"`
}

// Selection is authored by a user.
type Selection struct {
	ID string `json:"id"`
}

// Storage stores Users.
type Storage interface {
	CreateUser(authID string) (User, error)
	ReadUser(authID string) (User, error)
	ReadWritingAuthor(wri Writing) ([]User, error)
	ReadSelectionAuthor(sel Selection) ([]User, error)
}

// Service manages Users.
type Service interface {
	Storage
}

type service struct {
	Storage
}

// NewService manages users.
func NewService(store Storage) Service {
	return &service{store}
}

// Score shows how well a user's writings have been.
type Score struct {
	TotalPoints int `json:"total_points"`
}
