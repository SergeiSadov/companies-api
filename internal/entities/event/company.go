package event

type CompanyEvent struct {
	ID      string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Code    string `json:"code,omitempty"`
	Country string `json:"country,omitempty"`
	Website string `json:"website,omitempty"`
	Phone   string `json:"phone,omitempty"`
}

type IDEvent struct {
	ID string `json:"id,omitempty"`
}
