package ipapi

import "errors"

var (
	ErrToManyRequests = errors.New("too many requests to ipapi")
)
