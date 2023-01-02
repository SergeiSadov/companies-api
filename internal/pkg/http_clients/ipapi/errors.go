package ipapi

import "errors"

var (
	ErrEmptyCodeInResponse = errors.New("empty code in response")
	ErrToManyRequests      = errors.New("too many requests to ipapi")
)
