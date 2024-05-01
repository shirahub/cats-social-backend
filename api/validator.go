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
