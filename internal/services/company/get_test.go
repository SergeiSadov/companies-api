package company

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func (s *ServiceTestSuite) Test_SuccessGetCompany() {
	expectedData := s.getExpectedCompanyData()
	addr := gofakeit.IPv4Address()

	body, err := json.Marshal(expectedData.apiCompany)
	s.NoError(err)

	s.companyRepo.EXPECT().Get(gomock.Any(), expectedData.apiCompany.ID).Times(1).Return(expectedData.repoCompany, nil)
	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

	req, err := http.NewRequest(http.MethodGet, s.server.URL+"/companies/"+strconv.Itoa(expectedData.apiCompany.ID), bytes.NewReader(body))
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, s.token)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)
	var resultCompany api.GetCompanyResponse
	err = json.Unmarshal(respBody, &resultCompany)
	s.NoError(err)

	s.Equal(expectedData.apiCompany.ID, resultCompany.ID)
	s.Equal(expectedData.apiCompany.Name, resultCompany.Name)
	s.Equal(expectedData.apiCompany.Code, resultCompany.Code)
	s.Equal(expectedData.apiCompany.Country, resultCompany.Country)
	s.Equal(expectedData.apiCompany.Phone, resultCompany.Phone)
	s.Equal(expectedData.apiCompany.Website, resultCompany.Website)
}
