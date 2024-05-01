package api

import (
	"regexp"
	"github.com/go-playground/validator/v10"
)

const pattern = `^([<>]=?|=)\s*(\d+)$`

func validateComparison(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch(fl.Field().String())

	if len(matches) != 3 {
		return false
	}

	return true
}

var races = map[string]struct{}{
	"Persian":    struct{}{},
	"Maine Coon": struct{}{},
	"Siamese":    struct{}{},
	"Ragdoll":    struct{}{},
	"Bengal":     struct{}{},
	"Sphynx":     struct{}{},
	"British Shorthair": struct{}{},
	"Abyssinian":    struct{}{},
	"Scottish Fold": struct{}{},
	"Birman":        struct{}{},
}

func validateRace(fl validator.FieldLevel) bool {
	_, exists := races[fl.Field().String()]
	return exists
}