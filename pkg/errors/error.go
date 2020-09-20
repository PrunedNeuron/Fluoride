package errors

import "errors"

// ErrMarshal is the error thrown when marshalling fails
var ErrMarshal = errors.New("Could not marshal struct into JSON")
