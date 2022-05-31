package gitlab

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/B1scuit/gitlab/auth"
	"github.com/B1scuit/gitlab/project"
	"github.com/B1scuit/gitlab/types"
)

type Client struct {
	baseURL    *url.URL
	httpClient types.HttpClientInterface

	Auth    *auth.Auth
	Project *project.Project
	Group   *Group
}

func NewDefaultClient() *Client {
	return NewClient(nil)
}

func NewClient(opts *types.ClientOptions) *Client {
	if opts == nil {
		opts = &types.ClientOptions{}
	}

	if opts.BaseURL == nil {
		u, _ := url.Parse("https://gitlab.com/api/v4")
		opts.BaseURL = u
	}

	if opts.HttpClient == nil {
		opts.HttpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}

	// If log is nil, have a log that goes nowhere
	// but avoids nill pointer derefs
	var osNull *os.File
	if opts.Log == nil || opts.ErrLog == nil {
		var err error
		osNull, err = os.Open("/dev/null")
		if err != nil {
			panic(err)
		}
	}

	if opts.Log == nil {
		opts.Log = log.New(osNull, "", 0)
	}
	if opts.ErrLog == nil {
		opts.ErrLog = log.New(osNull, "", 0)
	}

	return &Client{
		httpClient: opts.HttpClient,
		baseURL:    opts.BaseURL,

		Auth:    auth.NewAuth(opts),
		Project: project.NewProject(opts),
		Group:   &Group{},
	}
}
