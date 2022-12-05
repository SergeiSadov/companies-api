package validators

import (
	"companies-api/internal/entities/api"
)

type IValidator interface {
	ValidateCreateRequest(req *api.CreateCompanyRequest) (err error)
	ValidateGetRequest(req *api.GetCompanyRequest) (err error)
	ValidateUpdateRequest(req *api.UpdateCompanyRequest) (err error)
	ValidateDeleteRequest(req *api.DeleteCompanyRequest) (err error)
}

type Validators struct {
	CreateValidators []ValidatorFunc
	UpdateValidators []ValidatorFunc
	GetValidators    []IDValidatorFunc
	DeleteValidators []IDValidatorFunc
}

func (v *Validators) ValidateCreateRequest(req *api.CreateCompanyRequest) (err error) {
	for _, v := range v.CreateValidators {
		if validationErr := v(req.Company); validationErr != nil {
			return validationErr
		}
	}

	return nil
}

func (v *Validators) ValidateGetRequest(req *api.GetCompanyRequest) (err error) {
	for _, v := range v.GetValidators {
		if validationErr := v(req.ID); validationErr != nil {
			return validationErr
		}
	}

	return nil
}

func (v *Validators) ValidateUpdateRequest(req *api.UpdateCompanyRequest) (err error) {
	for _, v := range v.UpdateValidators {
		if validationErr := v(req.Company); validationErr != nil {
			return validationErr
		}
	}

	return nil
}

func (v *Validators) ValidateDeleteRequest(req *api.DeleteCompanyRequest) (err error) {
	for _, v := range v.DeleteValidators {
		if validationErr := v(req.ID); validationErr != nil {
			return validationErr
		}
	}

	return nil
}
