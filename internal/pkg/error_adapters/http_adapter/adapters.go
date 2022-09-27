package http_adapter

import (
	"errors"
	"net/http"

	internalerrs "companies-api/internal/pkg/errors"
)

type ErrorToHttpCodeAdapter func(err error) (code int)

func AdaptBadRequestError(err error) (code int) {
	code = getCode(errors.Unwrap(err))
	if code > 0 {
		return
	}

	code = getCode(err)

	return code
}

func getCode(err error) (code int) {
	switch err {
	case
		internalerrs.ErrBadRequest,
		internalerrs.ErrAlreadyExist,
		internalerrs.ErrInvalidCompanyName,
		internalerrs.ErrInvalidCompanyCode,
		internalerrs.ErrInvalidCompanyCountry,
		internalerrs.ErrInvalidCompanyWebsite,
		internalerrs.ErrInvalidCompanyPhone,
		internalerrs.ErrInvalidCompanyID:
		code = http.StatusBadRequest
	case internalerrs.ErrNotFound:
		code = http.StatusNotFound
	}

	return code
}
