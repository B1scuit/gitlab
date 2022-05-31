package auth

import (
	"context"

	"github.com/B1scuit/gitlab/errors"
	"github.com/B1scuit/gitlab/types"
)

func newPersonalTokenFunc(opts *types.ClientOptions) PersonalToken {
	return func(personalToken string, o ...func(*PersonalTokenOptions)) error {
		r := PersonalTokenOptions{}

		for _, v := range o {
			v(&r)
		}

		return r.Do(r.ctx, personalToken, opts)
	}
}

type PersonalToken func(string, ...func(*PersonalTokenOptions)) error

type PersonalTokenOptions struct {
	ctx context.Context
}

func (pto *PersonalTokenOptions) Do(ctx context.Context, token string, opts *types.ClientOptions) error {

	if token == "" {
		return errors.ErrPersonalTokenMissing
	}

	opts.PersonalToken = token

	return nil
}
