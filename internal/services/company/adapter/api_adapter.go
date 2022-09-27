package adapter

import (
	"encoding/json"

	"companies-api/internal/entities/api"
	"companies-api/internal/entities/event"
	"companies-api/internal/entities/repository"

	"github.com/segmentio/kafka-go"
)

type IAdapter interface {
	AdaptCreateReqToRepo(req *api.CreateCompanyRequest) (adapted *repository.Company)
	AdaptRepoToCreateResp(req *repository.Company) (adapted *api.CreateCompanyResponse)

	AdaptGetReqToRepo(req *api.GetCompanyRequest) (id int)
	AdaptRepoToGetResp(req *repository.Company) (adapted *api.GetCompanyResponse)

	AdaptUpdateReqToRepo(req *api.UpdateCompanyRequest) (adapted *repository.Company)
	AdaptRepoToUpdateResp(req *repository.Company) (adapted *api.UpdateCompanyResponse)

	AdaptDeleteReqToRepo(req *api.DeleteCompanyRequest) (id int)

	AdaptListReqToRepo(req *api.ListCompanyRequest) (adapted *repository.ListCompanyParams)
	AdaptRepoToListResp(req []repository.Company) (adapted []api.Company)

	AdaptCompanyRepoToKafka(req *repository.Company) (adapted kafka.Message, err error)
	AdaptIDEventToKafka(id int) (adapted kafka.Message, err error)
}

type Adapter struct {
}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (a Adapter) AdaptCreateReqToRepo(req *api.CreateCompanyRequest) (adapted *repository.Company) {
	return &repository.Company{
		ID:      req.ID,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	}
}

func (a Adapter) AdaptRepoToCreateResp(req *repository.Company) (adapted *api.CreateCompanyResponse) {
	return &api.CreateCompanyResponse{
		Company: api.Company{
			ID:      req.ID,
			Name:    req.Name,
			Code:    req.Code,
			Country: req.Country,
			Website: req.Website,
			Phone:   req.Phone,
		},
	}
}

func (a Adapter) AdaptGetReqToRepo(req *api.GetCompanyRequest) (id int) {
	return req.ID
}

func (a Adapter) AdaptRepoToGetResp(req *repository.Company) (adapted *api.GetCompanyResponse) {
	return &api.GetCompanyResponse{
		Company: api.Company{
			ID:      req.ID,
			Name:    req.Name,
			Code:    req.Code,
			Country: req.Country,
			Website: req.Website,
			Phone:   req.Phone,
		},
	}
}

func (a Adapter) AdaptUpdateReqToRepo(req *api.UpdateCompanyRequest) (adapted *repository.Company) {
	return &repository.Company{
		ID:      req.ID,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	}
}

func (a Adapter) AdaptRepoToUpdateResp(req *repository.Company) (adapted *api.UpdateCompanyResponse) {
	return &api.UpdateCompanyResponse{
		Company: api.Company{
			ID:      req.ID,
			Name:    req.Name,
			Code:    req.Code,
			Country: req.Country,
			Website: req.Website,
			Phone:   req.Phone,
		},
	}
}

func (a Adapter) AdaptDeleteReqToRepo(req *api.DeleteCompanyRequest) (id int) {
	return req.ID
}

func (a Adapter) AdaptListReqToRepo(req *api.ListCompanyRequest) (adapted *repository.ListCompanyParams) {
	return &repository.ListCompanyParams{
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
		Page:    req.Page,
		Size:    req.Size,
	}
}

func (a Adapter) AdaptRepoToListResp(req []repository.Company) (adapted []api.Company) {
	adapted = make([]api.Company, len(req))
	for i := range req {
		adapted[i].ID = req[i].ID
		adapted[i].Name = req[i].Name
		adapted[i].Code = req[i].Code
		adapted[i].Country = req[i].Country
		adapted[i].Website = req[i].Website
		adapted[i].Phone = req[i].Phone
	}

	return
}

func (a Adapter) AdaptCompanyRepoToKafka(req *repository.Company) (adapted kafka.Message, err error) {
	e := event.CompanyEvent{
		ID:      req.ID,
		Name:    req.Name,
		Code:    req.Code,
		Country: req.Country,
		Website: req.Website,
		Phone:   req.Phone,
	}
	data, err := json.Marshal(e)
	if err != nil {
		return adapted, err
	}

	return kafka.Message{
		Value: data,
	}, nil
}

func (a Adapter) AdaptIDEventToKafka(id int) (adapted kafka.Message, err error) {
	e := event.IDEvent{
		ID: id,
	}
	data, err := json.Marshal(e)
	if err != nil {
		return adapted, err
	}

	return kafka.Message{
		Value: data,
	}, nil
}
