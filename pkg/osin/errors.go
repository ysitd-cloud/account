package osin

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	AssertStringError = errors.New("could not assert to string")
)
