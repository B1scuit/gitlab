package project

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"
	"time"

	"github.com/B1scuit/gitlab/errors"
	"github.com/B1scuit/gitlab/types"
)

func newListFunc(opts *types.ClientOptions) List {
	return func(o ...func(*ListOptions)) (*types.ProjectList, error) {
		r := ListOptions{}

		for _, v := range o {
			v(&r)
		}

		return r.Do(r.ctx, opts)
	}
}

type List func(...func(*ListOptions)) (*types.ProjectList, error)

type ListOptions struct {
	ctx context.Context

	Archived bool

	IdAfter  int
	IdBefore int

	Imported           bool
	LastActivityAfter  time.Time
	LastActivityBefore time.Time

	Membership bool

	MinAccessLevel int
	OrderBy        string
	Pagnation      string
	PerPage        int

	Owned bool

	RepositoryChecksumFailed bool
	RepositoryStorage        string

	SearchNamespaces bool
	Search           string
	Simple           bool

	Sort       string
	Starred    bool
	Statistics bool
	Topic      []string

	Visibility *types.ProjectVisibility

	WikiChecksumFailed   bool
	WithCustomAttributes bool

	WithIssuesEnabled        bool
	WithMergeRequestsEnabled bool
	WithProgrammingLanguage  string
}

func (lo *ListOptions) Do(ctx context.Context, opts *types.ClientOptions) (*types.ProjectList, error) {
	u := opts.BaseURL
	u.Path = path.Join(u.Path, "projects")

	q := u.Query()

	if lo.Archived {
		q.Set("archived", "true")
	}

	if lo.IdAfter != 0 {
		q.Set("id_after", fmt.Sprint(lo.IdAfter))
	}

	if lo.IdBefore != 0 {
		q.Set("id_before", fmt.Sprint(lo.IdBefore))
	}

	if lo.Imported {
		q.Set("imported", "true")
	}

	if !lo.LastActivityAfter.IsZero() {
		q.Set("last_activity_after", lo.LastActivityAfter.Format(time.RFC3339))
	}

	if !lo.LastActivityBefore.IsZero() {
		q.Set("last_activity_before", lo.LastActivityBefore.Format(time.RFC3339))
	}

	if lo.Membership {
		q.Set("membership", "true")
	}

	if lo.MinAccessLevel != 0 {
		q.Set("min_access_level", fmt.Sprint(lo.MinAccessLevel))
	}

	if lo.OrderBy != "" {
		q.Set("order_by", lo.OrderBy)
	}

	if lo.Owned {
		q.Set("owned", "true")
	}

	if lo.RepositoryChecksumFailed {
		q.Set("repository_checksum_failed", "true")
	}

	if lo.RepositoryStorage != "" {
		q.Set("repository_storage", lo.RepositoryStorage)
	}

	u.RawQuery = q.Encode()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	if err := opts.AuthHeaders(&request.Header); err != nil {
		// As projects can operate without authentication, only warn
		// whebn auth is missing
		opts.ErrLog.Println(err)
	}

	response, err := opts.HttpClient.Do(request)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.ErrNoAuthentication
	}

	defer response.Body.Close()

	responseData := &types.ProjectList{}
	if err := json.NewDecoder(response.Body).Decode(responseData); err != nil {
		return nil, err
	}

	log.Println("Request succeeded")

	return responseData, nil
}

func (lo *List) SetContext(ctx context.Context) func(*ListOptions) {
	return func(lo *ListOptions) {
		lo.ctx = ctx
	}
}

func (lo *List) SetOwned(in bool) func(*ListOptions) {
	return func(lo *ListOptions) {
		lo.Owned = in
	}
}
