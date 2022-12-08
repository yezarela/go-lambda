package domain

import (
	"errors"
)

var (
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound = errors.New("Your requested Item is not found")
	// ErrConflict will throw if the current item already exists
	ErrConflict = errors.New("Item already exist")
	// ErrInvalidParameters will throw if the given request-body or params is not valid
	ErrInvalidParameters = errors.New("Parameters is not valid")
)
