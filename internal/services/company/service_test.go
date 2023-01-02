package company

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"companies-api/internal/entities/api"
	"companies-api/internal/entities/event"
	"companies-api/internal/entities/repository"
	"companies-api/internal/pkg/error_adapters/http_adapter"
	"companies-api/internal/pkg/http_clients/ipapi"
	"companies-api/internal/pkg/middlewares/auth"
	"companies-api/internal/pkg/middlewares/ip"
	"companies-api/internal/repositories/company"
	"companies-api/internal/services/company/adapter"
	"companies-api/internal/services/company/validators"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

func TestService(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}

type ServiceTestSuite struct {
	suite.Suite
	ctrl    *gomock.Controller
	service IService

	companyRepo *company.MockIRepository
	adapter     adapter.IAdapter
	ipapiClient *ipapi.MockIClient
	server      *httptest.Server
	token       string
}

func (s *ServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.companyRepo = company.NewMockIRepository(s.ctrl)
	s.adapter = adapter.NewAdapter()
	s.ipapiClient = ipapi.NewMockIClient(s.ctrl)

	s.service = NewService(
		s.companyRepo,
		s.adapter,
	)

	logger := zap.NewNop()

	authMW := auth.NewMiddleware("secret", logger)
	ipMW := ip.NewMiddleware(s.ipapiClient, logger, ip.NewErrorAdapter(ip.PreparedErrorMapping), ip.DefaultAllowedCountry)

	r := mux.NewRouter()
	SetRoutes(s.service, validators.PreparedValidators, http_adapter.New(http.StatusInternalServerError, http_adapter.AdaptBadRequestError), r, logger, authMW, ipMW)
	s.server = httptest.NewServer(r)
	claims := &auth.Claims{
		Username: "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	s.NoError(err)
	if err != nil {
		return
	}
	s.token = tokenString
}

func (s *ServiceTestSuite) TearDownSuite() {
	s.server.Close()
}

type expectedCompanyData struct {
	apiCompany          *api.Company
	repoCompany         *repository.Company
	eventCompany        *event.CompanyEvent
	kafkaMessageCompany kafka.Message
}

func (s *ServiceTestSuite) getExpectedCompanyData() *expectedCompanyData {
	testCompany := s.getCompany()
	repoCompany := &repository.Company{
		ID:      testCompany.ID,
		Name:    testCompany.Name,
		Code:    testCompany.Code,
		Country: testCompany.Country,
		Website: testCompany.Website,
		Phone:   testCompany.Phone,
	}
	kafkaEvent := &event.CompanyEvent{
		ID:      testCompany.ID,
		Name:    testCompany.Name,
		Code:    testCompany.Code,
		Country: testCompany.Country,
		Website: testCompany.Website,
		Phone:   testCompany.Phone,
	}
	data, err := json.Marshal(kafkaEvent)
	s.NoError(err)
	kafkaMsg := kafka.Message{Value: data}

	return &expectedCompanyData{
		apiCompany:          testCompany,
		repoCompany:         repoCompany,
		eventCompany:        kafkaEvent,
		kafkaMessageCompany: kafkaMsg,
	}
}

func (s *ServiceTestSuite) getCompany() *api.Company {
	return &api.Company{
		ID:      gofakeit.UUID(),
		Name:    gofakeit.Company(),
		Code:    gofakeit.UUID(),
		Country: gofakeit.Country(),
		Website: gofakeit.URL(),
		Phone:   gofakeit.Phone(),
	}
}
