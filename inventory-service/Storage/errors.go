package Storage

import "errors"

var (
	ErrAlreadyExists = errors.New("object already exists")
	ErrNotFound      = errors.New("object not found")
)
