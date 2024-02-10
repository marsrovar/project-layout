package errs

import "errors"

var (
	ErrInputInvalidParameter = errors.New("InputInvalidParameter")
	ErrSystem                = errors.New("System")
	ErrUnauthorized          = errors.New("Unauthorized")
)
