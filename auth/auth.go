package auth

import "github.com/B1scuit/gitlab/types"

type Auth struct {
	PersonalToken PersonalToken
}

func NewAuth(opts *types.ClientOptions) *Auth {
	return &Auth{
		PersonalToken: newPersonalTokenFunc(opts),
	}
}
