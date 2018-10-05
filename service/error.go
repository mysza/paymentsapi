package service

// InputError indicates that there was something wrong with the input.
type InputError struct {
	message string
}

func (e *InputError) Error() string {
	return e.message
}

// NewInputError creates a new InputError instance
func NewInputError(message string) *InputError {
	return &InputError{message: message}
}

// NotFoundError indicates that the request failed because element was not found.
type NotFoundError struct {
	message string
}

func (e *NotFoundError) Error() string {
	return e.message
}

// NewNotFoundError creates a new NotFoundError.
func NewNotFoundError(message string) *NotFoundError {
	return &NotFoundError{message: message}
}
