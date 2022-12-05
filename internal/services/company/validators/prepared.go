package validators

var (
	PreparedValidators = &Validators{
		CreateValidators: []ValidatorFunc{
			ValidateName, ValidateCode, ValidateCountry, ValidateWebsite, ValidatePhone,
		},
		UpdateValidators: []ValidatorFunc{
			ValidateName, ValidateCode, ValidateCountry, ValidateWebsite, ValidatePhone,
		},
		GetValidators: []IDValidatorFunc{
			ValidateID,
		},
		DeleteValidators: []IDValidatorFunc{
			ValidateID,
		},
	}
)
