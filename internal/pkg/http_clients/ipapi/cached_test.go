package ipapi

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	lru "github.com/hashicorp/golang-lru"
	"github.com/stretchr/testify/assert"
)

func TestCachedClient_GetCountryCode_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHttpClient := NewMockHttpClient(ctrl)

	mockHttpClient.EXPECT().Get(fmt.Sprintf(UrlTpl, testIpOk)).Return(&http.Response{
		Body:       io.NopCloser(strings.NewReader(ipApiOk)),
		StatusCode: http.StatusOK,
	}, nil).
		//single call expected as the result must be placed in cache
		Times(1)
	client := NewClient(mockHttpClient, NewErrAdapter(map[int]error{
		StatusTooManyRequests: ErrToManyRequests,
		StatusRateLimited:     ErrToManyRequests,
	}))

	cache, err := lru.New(32)
	assert.NoError(t, err)

	cachedClient := NewCachedClient(client, cache)

	t.Run("cache miss", func(t *testing.T) {
		gotCode, err := cachedClient.GetCountryCode(testIpOk)
		assert.NoError(t, err)
		assert.Equal(t, testCountryCodeOk, gotCode)
	})

	t.Run("cached response", func(t *testing.T) {
		gotCode, err := cachedClient.GetCountryCode(testIpOk)
		assert.NoError(t, err)
		assert.Equal(t, testCountryCodeOk, gotCode)
	})
}

func TestCachedClient_GetCountryCode_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockHttpClient := NewMockHttpClient(ctrl)
	mockHttpClient.EXPECT().Get(fmt.Sprintf(UrlTpl, testIpFail)).Return(&http.Response{
		StatusCode: StatusTooManyRequests,
	}, errors.New("error"))
	client := NewClient(mockHttpClient, NewErrAdapter(map[int]error{
		StatusTooManyRequests: ErrToManyRequests,
	}))
	cache, err := lru.New(32)
	assert.NoError(t, err)

	cachedClient := NewCachedClient(client, cache)

	_, err = cachedClient.GetCountryCode(testIpFail)
	assert.Error(t, err)
}
