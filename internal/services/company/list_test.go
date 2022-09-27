package company

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"companies-api/internal/entities/api"
	"companies-api/internal/entities/repository"
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func (s *ServiceTestSuite) Test_SuccessListCompanies() {
	expectedData := s.getExpectedCompanyData()
	addr := gofakeit.IPv4Address()

	queryParams := url.Values{}
	queryParams.Add("name", expectedData.apiCompany.Name)
	queryParams.Add("code", expectedData.apiCompany.Code)
	queryParams.Add("country", expectedData.apiCompany.Country)
	queryParams.Add("website", expectedData.apiCompany.Website)
	queryParams.Add("phone", expectedData.apiCompany.Phone)
	queryParams.Add("page", "1")
	queryParams.Add("size", "1")

	s.companyRepo.EXPECT().List(gomock.Any(), &repository.ListCompanyParams{
		Name:    expectedData.apiCompany.Name,
		Code:    expectedData.apiCompany.Code,
		Country: expectedData.apiCompany.Country,
		Website: expectedData.apiCompany.Website,
		Phone:   expectedData.apiCompany.Phone,
		Page:    1,
		Size:    1,
	}).Times(1).Return([]repository.Company{*expectedData.repoCompany}, nil)
	s.companyRepo.EXPECT().Count(gomock.Any(), &repository.ListCompanyParams{
		Name:    expectedData.apiCompany.Name,
		Code:    expectedData.apiCompany.Code,
		Country: expectedData.apiCompany.Country,
		Website: expectedData.apiCompany.Website,
		Phone:   expectedData.apiCompany.Phone,
		Page:    1,
		Size:    1,
	}).Times(1).Return(1, nil)
	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/companies?"+queryParams.Encode(), nil)
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, s.token)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)
	var resultCompany api.ListCompanyResponse
	err = json.Unmarshal(respBody, &resultCompany)
	s.NoError(err)

	s.Equal(expectedData.apiCompany.ID, resultCompany.Data[0].ID)
	s.Equal(expectedData.apiCompany.Name, resultCompany.Data[0].Name)
	s.Equal(expectedData.apiCompany.Code, resultCompany.Data[0].Code)
	s.Equal(expectedData.apiCompany.Country, resultCompany.Data[0].Country)
	s.Equal(expectedData.apiCompany.Phone, resultCompany.Data[0].Phone)
	s.Equal(expectedData.apiCompany.Website, resultCompany.Data[0].Website)
}
