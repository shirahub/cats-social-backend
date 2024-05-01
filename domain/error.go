package domain

import "errors"

var (
	ErrNotFound = errors.New("resource not found")
	ErrMatchResponded = errors.New("match has already been responded")
)