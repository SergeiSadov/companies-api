package ip

import "companies-api/internal/pkg/http_clients/ipapi"

var (
	PreparedErrorMapping = map[error]error{
		//bypass too many requests error
		ipapi.ErrToManyRequests: nil,
	}
)

// IErrorAdapter is used to adapt and bypass client errors
type IErrorAdapter interface {
	AdaptError(initalErr error) (err error)
}

type ErrorAdapter struct {
	errorMapping map[error]error
}

func NewErrorAdapter(errorMapping map[error]error) *ErrorAdapter {
	return &ErrorAdapter{errorMapping: errorMapping}
}

func (a *ErrorAdapter) AdaptError(initalErr error) (err error) {
	foundErr, found := a.errorMapping[initalErr]
	if found {
		return foundErr
	}

	return initalErr
}
