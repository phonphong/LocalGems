package errors

type NotFoundError struct {
	message string
}

func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{
		message: message,
	}
}

func (e *NotFoundError) Error() string {
	return e.message
}

type ValidationError struct {
	message string
}

func NewValidationError(message string) *ValidationError {
	return &ValidationError{
		message: message,
	}
}

func (e *ValidationError) Error() string {
	return e.message
}

