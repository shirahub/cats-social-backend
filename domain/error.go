package domain

import "errors"

var (
	ErrNotFound = errors.New("resource not found")
	ErrEmailTaken = errors.New("email has been taken")
	ErrCatInMatch = errors.New("sex is not editable due to cat participating in one or more matches")
	ErrMatchResponded = errors.New("match has already been responded")
	ErrMatchWithOwnedCat = errors.New("cannot match with owned cat")
	ErrMatchWithSameSex = errors.New("cannot match with same sex")
	ErrMatchExists = errors.New("cats already participated in an existing match")
	ErrMatchWithTaken = errors.New("either issuer or receiver cat has been matched with another cat")
)