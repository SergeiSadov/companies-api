package company

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func (s *ServiceTestSuite) Test_SuccessCreateCompany() {
	expectedData := s.getExpectedCompanyData()
	addr := gofakeit.IPv4Address()

	body, err := json.Marshal(expectedData.apiCompany)
	s.NoError(err)

	s.companyRepo.EXPECT().Create(gomock.Any(), expectedData.repoCompany).Times(1).Return(expectedData.repoCompany, nil)
	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/companies", bytes.NewReader(body))
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, s.token)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)

	respBody, err := io.ReadAll(resp.Body)
	s.NoError(err)
	var resultCompany api.CreateCompanyResponse
	err = json.Unmarshal(respBody, &resultCompany)
	s.NoError(err)

	s.Equal(expectedData.apiCompany.ID, resultCompany.ID)
	s.Equal(expectedData.apiCompany.Name, resultCompany.Name)
	s.Equal(expectedData.apiCompany.Code, resultCompany.Code)
	s.Equal(expectedData.apiCompany.Country, resultCompany.Country)
	s.Equal(expectedData.apiCompany.Phone, resultCompany.Phone)
	s.Equal(expectedData.apiCompany.Website, resultCompany.Website)
}

func (s *ServiceTestSuite) Test_FailCreateCompanyWrongIP() {
	expectedData := s.getExpectedCompanyData()
	addr := gofakeit.IPv4Address()

	body, err := json.Marshal(expectedData.apiCompany)
	s.NoError(err)

	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return("UK", nil)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/companies", bytes.NewReader(body))
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, s.token)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)
	s.Equal(http.StatusForbidden, resp.StatusCode)
}

func (s *ServiceTestSuite) Test_FailCreateCompanyWrongToken() {
	expectedData := s.getExpectedCompanyData()
	addr := gofakeit.IPv4Address()

	body, err := json.Marshal(expectedData.apiCompany)
	s.NoError(err)

	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/companies", bytes.NewReader(body))
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, gofakeit.UUID())

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)
	s.Equal(http.StatusUnauthorized, resp.StatusCode)
}

func (s *ServiceTestSuite) Test_FailCreateCompanyValidationErr() {
	expectedData := []api.Company{
		{Name: ""},
		{Name: "Name", Country: "TEST"},
		{Name: "Name", Country: "US", Phone: "QWERTY"},
	}
	addr := gofakeit.IPv4Address()

	for _, company := range expectedData {
		body, err := json.Marshal(company)
		s.NoError(err)

		s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

		req, err := http.NewRequest(http.MethodPost, s.server.URL+"/companies", bytes.NewReader(body))
		s.NoError(err)
		req.Header.Add("X-Forwarded-For", addr)
		req.Header.Add(auth.Authorization, s.token)

		resp, err := http.DefaultClient.Do(req)
		s.NoError(err)
		s.Equal(http.StatusBadRequest, resp.StatusCode)
	}

}
