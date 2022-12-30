package api

type Company struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
	Website string `json:"website,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type CreateCompanyRequest struct {
	Company
}

type CreateCompanyResponse struct {
	Company
}

type UpdateCompanyRequest struct {
	Company
}

type UpdateCompanyResponse struct {
	Company
}

type ListCompanyRequest struct {
	Name    string
	Code    string
	Country string
	Website string
	Phone   string
	Page    int
	Size    int
}

type CompanyData struct {
	Company
}

type ListCompanyResponse struct {
	Data []Company `json:"data"`
	Meta Meta      `json:"meta"`
}

type GetCompanyRequest struct {
	ID string
}

type GetCompanyResponse struct {
	Company
}

type DeleteCompanyRequest struct {
	ID string
}

type CompanyValidationParams struct {
	NameMaxLen    int
	CodeMaxLen    int
	CountryMaxLen int
	WebsiteMaxLen int
	PhoneMaxLen   int
	PhoneMinLen   int
}
