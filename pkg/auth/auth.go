package auth

// Service can authorize Users.
type Service interface {
	Signup(creds Creds) (token string, err error)
	Login(creds Creds) (token string, err error)
	Verify(token string) (username string, err error)
}

// Creds are used for signing up and logging in.
type Creds struct {
	Email string `json:"email"`
	Pwd   string `json:"password"`
}
