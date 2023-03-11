package servicefactory

// multiError help aggregates error messages.
type multiError struct {
	errors []string
}

// Error returns the aggregated errors.
func (multiError *multiError) Error() string {
	errorsLiteral := ""
	for _, e := range multiError.errors {
		errorsLiteral += e + "\n"
	}
	return errorsLiteral
}

// newErrors news multiError that allows input of multiple
// errors.
func newErrors(errors []string) error {
	return &multiError{
		errors: errors,
	}
}
