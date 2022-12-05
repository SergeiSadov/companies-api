package ipapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	testIpOk   = "192.168.8.1.1"
	testIpFail = "0.0.0.0"

	testCountryCodeOk = "CY"
)

var (
	ipApiOk = `{
		"ip": "` + testIpOk + `",
		"country_code": "` + testCountryCodeOk + `"
	}`
)

type mockHttpClientSetter func(c *MockHttpClient)

func TestClient_GetCountryCode_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHttpClient := NewMockHttpClient(ctrl)
	mockHttpClient.EXPECT().Get(fmt.Sprintf(UrlTpl, testIpOk)).Return(&http.Response{
		Body: io.NopCloser(strings.NewReader(ipApiOk)),
	}, nil)
	client := NewClient(mockHttpClient, NewErrAdapter(map[int]error{
		StatusTooManyRequests: ErrToManyRequests,
	}))
	gotCode, err := client.GetCountryCode(testIpOk)
	assert.NoError(t, err)
	assert.Equal(t, testCountryCodeOk, gotCode)
}

func TestClient_GetCountryCode_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHttpClient := NewMockHttpClient(ctrl)
	mockHttpClient.EXPECT().Get(fmt.Sprintf(UrlTpl, testIpFail)).Return(&http.Response{
		StatusCode: StatusTooManyRequests,
	}, errors.New("error"))
	client := NewClient(mockHttpClient, NewErrAdapter(map[int]error{
		StatusTooManyRequests: ErrToManyRequests,
	}))

	_, err := client.GetCountryCode(testIpFail)
	assert.Equal(t, ErrToManyRequests, err)
}
