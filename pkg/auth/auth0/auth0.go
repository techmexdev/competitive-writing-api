package auth0

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"errors"

	"github.com/techmexdev/competitive_writing_api/pkg/auth"
)

type auth0 struct {
	config      Config
	accessToken string
	url         string
}

// Config is auth0 configuration.
type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Audience     string `json:"audience"`
	DB           string // auth0 connection db name
}

// New uses Config to fetch access token from auth0 servers.
func New(config Config) (auth.Service, error) {
	url := "https://competitive-writing.auth0.com"
	a := auth0{config: config, url: url}
	at, err := a.getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("failed getting access token: %s", err)
	}
	a.accessToken = at

	return a, nil
}

// ErrAuth0 are errors Auth0 responds with in their response body.
type ErrAuth0 struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    string `json:"code"`
	Desc    struct {
		Rules []struct {
			Message  string `json:"message"`
			Code     string `json:"code"`
			Verified string `json:"verified"`
			Format   []int  `json:"format"`
			Items    []struct {
				Message  string `json:"message"`
				Code     string `json:"code"`
				Verified string `json:"verified"`
			}
		} `json:"rules"`
	} `json:"description"`
	StatusCode int    `json:"statusCode"`
	Policy     string `json:"policy"`
}

func (e ErrAuth0) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

type errJSON struct {
	Msg string `json:"error"`
}

func (e errJSON) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (a auth0) Signup(creds auth.Creds) (token string, err error) {
	type payload struct {
		ClientID   string `json:"client_id"`
		Email      string `json:"email"`
		Pwd        string `json:"password"`
		Connection string `json:"connection"`
	}

	pl := payload{
		ClientID:   a.config.ClientID,
		Email:      creds.Email,
		Pwd:        creds.Pwd,
		Connection: a.config.DB,
	}

	b, err := json.MarshalIndent(pl, "", " ")

	req, err := http.NewRequest(http.MethodPost, a.url+"/dbconnections/signup", bytes.NewBuffer(b))
	if err != nil {
		return "", fmt.Errorf("failed creating request: %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authentication", "Bearer "+a.accessToken)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to auth0: %s", err)
	}

	if res.StatusCode >= 400 {
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		res.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		var err ErrAuth0
		json.Unmarshal(body, &err)

		if err.Name == "" {
			var err errJSON
			json.Unmarshal(body, &err)
			return "", err
		}
		err.StatusCode = res.StatusCode
		return "", err
	}

	type body struct {
		ID string `json:"_id"`
	}

	var bd body
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&bd)

	return bd.ID, nil
}

func (a auth0) Login(creds auth.Creds) (token string, err error) {
	return "", errors.New("not implemented")
}

func (a auth0) Verify(token string) (username string, err error) {
	return "", errors.New("not implemented")
}

func (a auth0) getAccessToken() (accessToken string, err error) {
	type payload struct {
		Config
		GrantType string `json:"grant_type"`
	}

	pl := payload{Config: a.config, GrantType: "client_credentials"}
	b, err := json.MarshalIndent(pl, "", " ")

	req, _ := http.NewRequest(http.MethodPost, a.url+"/oauth/token", bytes.NewBuffer(b))

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request to auth0: %s", err)
	}

	type body struct {
		AccessToken string `json:"access_token"`
	}
	var bd body
	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&bd)

	return bd.AccessToken, nil
}
