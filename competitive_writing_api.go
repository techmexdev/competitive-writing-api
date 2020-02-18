package competitivewritingapi

// Auth can authorize Users.
type Auth interface {
	Verify(token string) (err error)
	CreateToken(username string) (token string, err error)
}

// Storage can store/retrieve resources.
type Storage struct {
	User      UserStorage
	Passage   PassageStorage
	Analysis  AnalysisStorage
	Selection SelectionStorage
}

// UserStorage stores Users.
type UserStorage interface {
	Create(usr User) (err error)
	Read(username string) (usr User, err error)
	List() (usrs []User, err error)
}

// PassageStorage stores Users.
type PassageStorage interface {
	Create(pass Passage) (err error)
	Read(id string) (pass Passage, err error)
	List() (passs []Passage, err error)
}

// AnalysisStorage stores Users.
type AnalysisStorage interface {
	Create(ana Analysis) (err error)
	Read(id string) (ana Analysis, err error)
	List() (anas []Analysis, err error)
}

// SelectionStorage stores Users.
type SelectionStorage interface {
	Create(sel Selection) (err error)
	Read(id string) (sel Selection, err error)
	List() (sels []Selection, err error)
}

// User is a user of the app.
type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Passage is text that users write a review of.
type Passage struct {
	ID     string `json:"id"`
	Author string `json:"author"`
	Book   string `json:"book"`
	Text   string `json:"text"`
}

// Analysis is a user's re-write/commentary on a passagraph.
type Analysis struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Passage  `json:"passage"`
	Text     string `json:"text"`
}

// Selection is a user's choice of best review out of reviews.
type Selection struct {
	ID              string `json:"id"`
	Username        string `json:"username"`
	Passage         `json:"passage"`
	BestAnalysis    Analysis   `json:"best_analysis"`
	AnalysisChoices []Analysis `json:"analysis_choices"`
}
