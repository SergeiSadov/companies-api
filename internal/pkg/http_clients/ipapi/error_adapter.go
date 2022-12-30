package ipapi

const (
	StatusTooManyRequests = 104
	StatusRateLimited     = 429
)

type IErrorAdapter interface {
	AdaptStatusToErr(code int, initialErr error) (err error)
}

type ErrorAdapter struct {
	codeToErrMapping map[int]error
}

func NewErrAdapter(codeToErrMapping map[int]error) *ErrorAdapter {
	return &ErrorAdapter{
		codeToErrMapping: codeToErrMapping,
	}
}

func (adapter ErrorAdapter) AdaptStatusToErr(code int, initialErr error) (err error) {
	foundErr, found := adapter.codeToErrMapping[code]
	if found {
		return foundErr
	}

	return initialErr
}
