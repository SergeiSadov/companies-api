package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/error_adapters/http_adapter"
	"companies-api/internal/pkg/middlewares/auth"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang-jwt/jwt/v4"
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
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

	secret string
	token  string
	server *httptest.Server
}

func (s *ServiceTestSuite) SetupTest() {
	s.ctrl = gomock.NewController(s.T())
	s.secret = "secret"

	s.service = NewService(s.secret)
	logger, err := zap.NewDevelopment()
	s.NoError(err)

	r := mux.NewRouter()
	SetRoutes(s.service, http_adapter.New(http.StatusInternalServerError, http_adapter.AdaptBadRequestError), r, logger)
	s.server = httptest.NewServer(r)
	claims := &auth.Claims{
		Username: "user",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secret))
	s.NoError(err)
	if err != nil {
		return
	}
	s.token = tokenString
}

func (s *ServiceTestSuite) TearDownSuite() {
	s.server.Close()
}

func (s *ServiceTestSuite) Test_SuccessAuth() {
	data := api.AuthRequest{
		Password: gofakeit.Password(false, false, false, false, false, 5),
		Username: gofakeit.Username(),
	}

	body, err := json.Marshal(data)
	s.NoError(err)

	req, err := http.NewRequest(http.MethodPost, s.server.URL+"/auth", bytes.NewReader(body))
	s.NoError(err)

	resp, err := http.DefaultClient.Do(req)
	s.NoError(err)

	respBody, err := ioutil.ReadAll(resp.Body)
	s.NoError(err)

	var authResponse api.AuthResponse
	fmt.Println(string(respBody))
	err = json.Unmarshal(respBody, &authResponse)
	s.NoError(err)

	s.NotEmpty(authResponse.Token)
}
