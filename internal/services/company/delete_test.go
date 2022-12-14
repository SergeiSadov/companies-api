package company

import (
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
)

func (s *ServiceTestSuite) Test_SuccessDeleteCompany() {
	addr := gofakeit.IPv4Address()
	id := gofakeit.UUID()

	s.companyRepo.EXPECT().Delete(gomock.Any(), id).Times(1).Return(nil)
	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

	req, err := http.NewRequest(http.MethodDelete, s.server.URL+"/companies/"+id, nil)
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, s.token)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
}
