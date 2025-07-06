package validators

import (
	"local-gems-server/internal/core/errors"
	"regexp"
	"strconv"
)

func ValidateEmail(email string) error {
	// Simple email validation
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !re.MatchString(email) {
		return errors.NewValidationError("invalid email format")
	}
	return nil
}

func ValidateLength(value string, min, max int) error {
	if len(value) < min || len(value) > max {
		return errors.NewValidationError("value length must be between " + strconv.Itoa(min) + " and " + strconv.Itoa(max))
	}
	return nil
}
