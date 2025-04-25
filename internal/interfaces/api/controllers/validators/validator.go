package validators

import (
	"localgems/internal/core/errors"
	"regexp"
)

// Cung cấp các helpers cho validation
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
		return errors.NewValidationError("value length must be between " + string(min) + " and " + string(max))
	}
	return nil
}
