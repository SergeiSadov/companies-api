package http_adapter

import (
	"fmt"
	"net/http"
	"testing"

	internalerrs "companies-api/internal/pkg/errors"

	"github.com/stretchr/testify/assert"
)

func TestErrorAdapter_AdaptToHttpCode(t *testing.T) {
	adapter := New(http.StatusInternalServerError, AdaptBadRequestError)

	tests := []struct {
		name        string
		a           *ErrorAdapter
		errs        []error
		wantAdapted int
	}{
		{
			"Bad request",
			adapter,
			[]error{
				internalerrs.ErrBadRequest,
				internalerrs.ErrAlreadyExist,
				internalerrs.ErrInvalidCompanyName,
				internalerrs.ErrInvalidCompanyCode,
				internalerrs.ErrInvalidCompanyCountry,
				internalerrs.ErrInvalidCompanyWebsite,
				internalerrs.ErrInvalidCompanyPhone,
				internalerrs.ErrInvalidCompanyID,

				//wrapped errors

				fmt.Errorf("%w", internalerrs.ErrBadRequest),
				fmt.Errorf("%w", internalerrs.ErrAlreadyExist),
				fmt.Errorf("%w", internalerrs.ErrInvalidCompanyName),
				fmt.Errorf("%w", internalerrs.ErrInvalidCompanyCode),
				fmt.Errorf("%w", internalerrs.ErrInvalidCompanyCountry),
				fmt.Errorf("%w", internalerrs.ErrInvalidCompanyWebsite),
				fmt.Errorf("%w", internalerrs.ErrInvalidCompanyPhone),
				fmt.Errorf("%w", internalerrs.ErrInvalidCompanyID),
			},
			http.StatusBadRequest,
		},
		{
			"Not found",
			adapter,
			[]error{
				internalerrs.ErrNotFound,

				//wrapped errors

				fmt.Errorf("%w", internalerrs.ErrNotFound),
			},
			http.StatusNotFound,
		},
	}
	for _, tt := range tests {
		for i := range tt.errs {
			t.Run(tt.name, func(t *testing.T) {
				assert.Equal(t, tt.wantAdapted, tt.a.AdaptToHttpCode(tt.errs[i]))
			})
		}
	}
}
