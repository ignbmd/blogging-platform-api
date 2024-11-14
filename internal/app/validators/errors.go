package validators

type ValidationError struct {
	Errors map[string]string
}

func (e *ValidationError) Error() string {
	return "validation failed"
}

func NewValidationError(errors map[string]string) *ValidationError {
	return &ValidationError{
		Errors: errors,
	}
}
