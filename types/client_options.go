package types

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/B1scuit/gitlab/errors"
)

type HttpClientInterface interface {
	Do(*http.Request) (*http.Response, error)
}

type ClientOptions struct {
	BaseURL    *url.URL
	HttpClient HttpClientInterface

	Log    *log.Logger
	ErrLog *log.Logger

	// Auth holders
	PersonalToken string
	BearerToken   string
}

func (co *ClientOptions) AuthHeaders(in *http.Header) error {

	// Authenticating using a personal token
	if co.PersonalToken != "" {
		in.Set("PRIVATE-TOKEN", co.PersonalToken)
		return nil
	}

	// Authentication using oAuth
	if co.BearerToken != "" {
		in.Set("Authorization", fmt.Sprintf("Bearer: %v", co.BearerToken))
		return nil
	}

	return errors.ErrNoAuthentication
}
