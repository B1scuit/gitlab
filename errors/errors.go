package errors

import err "errors"

var (
	// Auth Errors
	ErrPersonalTokenMissing = err.New("personal token is missing")
	ErrNoAuthentication     = err.New("no authentication strategy present")
)
