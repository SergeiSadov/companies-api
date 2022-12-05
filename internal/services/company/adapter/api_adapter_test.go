package adapter

import (
	"companies-api/internal/entities/api"
	"companies-api/internal/entities/event"
	"companies-api/internal/entities/repository"
	"encoding/json"
	"testing"

	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
)

const (
	testProjectID = 1
	testName      = "name"
	testCode      = "1234-test"
	testCountry   = "US"
	testWebsite   = "http://example.com"
	testPhone     = "12345677890"
	testPage      = 1
	testSize      = 5
)

var (
	companyRequest = api.Company{
		ID:      testProjectID,
		Name:    testName,
		Code:    testCode,
		Country: testCountry,
		Website: testWebsite,
		Phone:   testPhone,
	}

	companyRepository = repository.Company{
		ID:      companyRequest.ID,
		Name:    companyRequest.Name,
		Code:    companyRequest.Code,
		Country: companyRequest.Country,
		Website: companyRequest.Website,
		Phone:   companyRequest.Phone,
	}

	listRequest = api.ListCompanyRequest{
		Name:    testName,
		Code:    testCode,
		Country: testCountry,
		Website: testWebsite,
		Phone:   testPhone,
		Page:    testPage,
		Size:    testSize,
	}
)

func TestAdapter_AdaptCreateReqToRepo(t *testing.T) {
	adapter := NewAdapter()
	tests := []struct {
		name        string
		a           *Adapter
		req         *api.CreateCompanyRequest
		wantAdapted *repository.Company
	}{
		{
			"all fields are present",
			adapter,
			&api.CreateCompanyRequest{
				Company: companyRequest,
			},
			&companyRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptCreateReqToRepo(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptRepoToCreateResp(t *testing.T) {
	adapter := NewAdapter()
	tests := []struct {
		name        string
		a           *Adapter
		req         *repository.Company
		wantAdapted *api.CreateCompanyResponse
	}{
		{
			"all fields are present",
			adapter,
			&companyRepository,
			&api.CreateCompanyResponse{
				Company: companyRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptRepoToCreateResp(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptGetReqToRepo(t *testing.T) {
	adapter := NewAdapter()
	getCompanyRequest := api.GetCompanyRequest{
		ID: testProjectID,
	}

	tests := []struct {
		name   string
		a      *Adapter
		req    *api.GetCompanyRequest
		wantId int
	}{
		{
			"id is present",
			adapter,
			&getCompanyRequest,
			testProjectID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptGetReqToRepo(tt.req)
			assert.Equal(t, tt.wantId, actual)
		})
	}
}

func TestAdapter_AdaptRepoToGetResp(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name        string
		a           *Adapter
		req         *repository.Company
		wantAdapted *api.GetCompanyResponse
	}{
		{
			"all fields are present",
			adapter,
			&companyRepository,
			&api.GetCompanyResponse{
				Company: companyRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptRepoToGetResp(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptUpdateReqToRepo(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name        string
		a           *Adapter
		req         *api.UpdateCompanyRequest
		wantAdapted *repository.Company
	}{
		{
			"all fields are present",
			adapter,
			&api.UpdateCompanyRequest{
				Company: companyRequest,
			},
			&companyRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptUpdateReqToRepo(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptRepoToUpdateResp(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name        string
		a           *Adapter
		req         *repository.Company
		wantAdapted *api.UpdateCompanyResponse
	}{
		{
			"all fields are present",
			adapter,
			&companyRepository,
			&api.UpdateCompanyResponse{
				Company: companyRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptRepoToUpdateResp(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptDeleteReqToRepo(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name   string
		a      *Adapter
		req    *api.DeleteCompanyRequest
		wantId int
	}{
		{
			"id is present",
			adapter,
			&api.DeleteCompanyRequest{
				ID: testProjectID,
			},
			testProjectID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptDeleteReqToRepo(tt.req)
			assert.Equal(t, tt.wantId, actual)
		})
	}
}

func TestAdapter_AdaptListReqToRepo(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name        string
		a           *Adapter
		req         *api.ListCompanyRequest
		wantAdapted *repository.ListCompanyParams
	}{
		{
			"all fields are present",
			adapter,
			&listRequest,
			&repository.ListCompanyParams{
				Name:    listRequest.Name,
				Code:    listRequest.Code,
				Country: listRequest.Country,
				Website: listRequest.Website,
				Phone:   listRequest.Phone,
				Page:    listRequest.Page,
				Size:    listRequest.Size,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptListReqToRepo(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptRepoToListResp(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name        string
		a           *Adapter
		req         []repository.Company
		wantAdapted []api.Company
	}{
		{
			"all fields are present",
			adapter,
			[]repository.Company{
				companyRepository,
			},
			[]api.Company{
				companyRequest,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.a.AdaptRepoToListResp(tt.req)
			assert.Equal(t, tt.wantAdapted, actual)
		})
	}
}

func TestAdapter_AdaptCompanyRepoToKafka(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name string
		a    *Adapter
		req  *repository.Company
	}{
		{
			"all fields are present",
			adapter,
			&companyRepository,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := event.CompanyEvent{
				ID:      tt.req.ID,
				Name:    tt.req.Name,
				Code:    tt.req.Code,
				Country: tt.req.Country,
				Website: tt.req.Website,
				Phone:   tt.req.Phone,
			}
			expected, err := json.Marshal(e)
			assert.NoError(t, err)

			actual, err := tt.a.AdaptCompanyRepoToKafka(tt.req)
			assert.NoError(t, err)
			assert.Equal(t, kafka.Message{Value: expected}, actual)
		})
	}
}

func TestAdapter_AdaptIDEventToKafka(t *testing.T) {
	adapter := NewAdapter()

	tests := []struct {
		name string
		a    *Adapter
		id   int
	}{
		{
			"all fields are present",
			adapter,
			testProjectID,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := event.IDEvent{
				ID: tt.id,
			}
			expected, err := json.Marshal(e)
			assert.NoError(t, err)

			actual, err := tt.a.AdaptIDEventToKafka(tt.id)
			assert.NoError(t, err)
			assert.Equal(t, kafka.Message{Value: expected}, actual)
		})
	}
}
