package apperrors

import "errors"

var (
	ErrQueryDB   = errors.New("error loading DB")
	ErrScanDB    = errors.New("error scaning DB")
	ErrNotFound  = errors.New("not found")
	ErrEmptyList = errors.New("empty list")
)
