package validators

import (
	"companies-api/internal/entities/api"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidators_ValidateCreate(t *testing.T) {
	tests := []struct {
		name    string
		v       *Validators
		req     *api.CreateCompanyRequest
		wantErr bool
	}{
		{
			"Success",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   gofakeit.Phone(),
				},
			},
			false,
		},
		{
			"Error#EmptyCompany",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    "",
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#CompanyNameTooLong",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.LetterN(256),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#EmptyCode",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    "",
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#CodeTooLong",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.LetterN(256),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#CountryNotExists",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.LetterN(1),
					Website: gofakeit.URL(),
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#EmptyWebsite",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: "",
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#WebsiteTooLong",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.LetterN(256),
					Phone:   gofakeit.Phone(),
				},
			},
			true,
		},
		{
			"Error#PhoneNaN",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   "test",
				},
			},
			true,
		},
		{
			"Error#PhoneTooShort",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   "1",
				},
			},
			true,
		},
		{
			"Error#PhoneTooLong",
			PreparedValidators,
			&api.CreateCompanyRequest{
				api.Company{
					Name:    gofakeit.Company(),
					Code:    gofakeit.UUID(),
					Country: gofakeit.CountryAbr(),
					Website: gofakeit.URL(),
					Phone:   "123456789101112",
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.v.ValidateCreateRequest(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidators_ValidateGet(t *testing.T) {
	tests := []struct {
		name    string
		v       *Validators
		req     *api.GetCompanyRequest
		wantErr bool
	}{
		{
			"Success",
			PreparedValidators,
			&api.GetCompanyRequest{
				ID: gofakeit.UUID(),
			},
			false,
		},
		{
			"Error#EmptyID",
			PreparedValidators,
			&api.GetCompanyRequest{
				"",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.v.ValidateGetRequest(tt.req)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
