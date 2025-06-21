package apperrors

import "errors"

var (
	ErrLoadCSV         = errors.New("failed to Load CSV")
	ErrCountryNotFound = errors.New("country not found")
)
