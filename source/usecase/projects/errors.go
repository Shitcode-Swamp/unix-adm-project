package projects

import "errors"

var ErrInvalidPath = errors.New("env paths must start with ~/")
