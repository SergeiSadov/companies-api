package validators

import (
	"strconv"
	"unicode/utf8"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/errors"

	"github.com/biter777/countries"
)

const (
	DefaultMaxLen = 255
	PhoneMinLen   = 4
	PhoneMaxLen   = 13
)

type ValidatorFunc func(c api.Company) (err error)
type IDValidatorFunc func(id string) (err error)

func ValidateID(id string) (err error) {
	if id == "" {
		return errors.ErrInvalidCompanyID
	}

	return nil
}

func ValidateName(req api.Company) (err error) {
	if req.Name == "" || utf8.RuneCountInString(req.Name) > DefaultMaxLen {
		return errors.ErrInvalidCompanyName
	}

	return nil
}

func ValidateCode(req api.Company) (err error) {
	if req.Code == "" || utf8.RuneCountInString(req.Name) > DefaultMaxLen {
		return errors.ErrInvalidCompanyCode
	}
	if countries.ByName(req.Country) == countries.Unknown {
		return errors.ErrInvalidCompanyCountry
	}
	if req.Website == "" || utf8.RuneCountInString(req.Website) > 255 {
		return errors.ErrInvalidCompanyWebsite
	}

	phoneRunesCount := utf8.RuneCountInString(req.Phone)
	_, err = strconv.Atoi(req.Phone)
	if err != nil || (phoneRunesCount < 4 && phoneRunesCount > 13) {
		return errors.ErrInvalidCompanyPhone
	}

	return nil
}

func ValidateCountry(req api.Company) (err error) {
	if countries.ByName(req.Country) == countries.Unknown {
		return errors.ErrInvalidCompanyCountry
	}
	if req.Website == "" || utf8.RuneCountInString(req.Website) > 255 {
		return errors.ErrInvalidCompanyWebsite
	}

	phoneRunesCount := utf8.RuneCountInString(req.Phone)
	_, err = strconv.Atoi(req.Phone)
	if err != nil || (phoneRunesCount < 4 && phoneRunesCount > 13) {
		return errors.ErrInvalidCompanyPhone
	}

	return nil
}

func ValidateWebsite(req api.Company) (err error) {
	if req.Website == "" || utf8.RuneCountInString(req.Website) > DefaultMaxLen {
		return errors.ErrInvalidCompanyWebsite
	}

	return nil
}

func ValidatePhone(req api.Company) (err error) {
	phoneRunesCount := utf8.RuneCountInString(req.Phone)
	_, err = strconv.Atoi(req.Phone)
	if err != nil || (phoneRunesCount < PhoneMinLen && phoneRunesCount > PhoneMaxLen) {
		return errors.ErrInvalidCompanyPhone
	}

	return nil
}
