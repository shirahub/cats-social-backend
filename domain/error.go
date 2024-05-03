package domain

import "errors"

var (
	ErrNotFound = errors.New("resource not found")
	ErrEmailTaken = errors.New("email has been taken")
	ErrMatchResponded = errors.New("match has already been responded")
	ErrMatchWithOwnedCat = errors.New("cannot match with owned cat")
	ErrMatchWithSameSex = errors.New("cannot match with same sex")
	ErrMatchWithTaken = errors.New("either issuer or receiver cat has been matched with another cat")
)