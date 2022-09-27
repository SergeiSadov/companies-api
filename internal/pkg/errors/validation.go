package errors

import "errors"

var (
	ErrInvalidCompanyName    = errors.New("invalid company name")
	ErrInvalidCompanyCode    = errors.New("invalid company code")
	ErrInvalidCompanyCountry = errors.New("invalid company country")
	ErrInvalidCompanyWebsite = errors.New("invalid company website")
	ErrInvalidCompanyPhone   = errors.New("invalid company phone")
	ErrInvalidCompanyID      = errors.New("invalid company id")
)
