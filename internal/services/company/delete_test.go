package company

import (
	"companies-api/internal/entities/event"
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"
	"encoding/json"
	"net/http"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/segmentio/kafka-go"
)

func (s *ServiceTestSuite) Test_SuccessDeleteCompany() {
	addr := gofakeit.IPv4Address()
	id := gofakeit.UUID()
	kafkaEvent := &event.IDEvent{
		ID: id,
	}
	data, err := json.Marshal(kafkaEvent)
	s.NoError(err)
	kafkaMsg := kafka.Message{Value: data}

	s.companyRepo.EXPECT().Delete(gomock.Any(), id).Times(1).Return(nil)
	s.deleteWriter.EXPECT().WriteMessages(gomock.Any(), kafkaMsg).Times(1).Return(nil)
	s.ipapiClient.EXPECT().GetCountryCode(addr).Times(1).Return(ip.DefaultAllowedCountry, nil)

	req, err := http.NewRequest(http.MethodDelete, s.server.URL+"/companies/"+id, nil)
	s.NoError(err)
	req.Header.Add("X-Forwarded-For", addr)
	req.Header.Add(auth.Authorization, s.token)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)
	s.Equal(http.StatusOK, resp.StatusCode)
}
