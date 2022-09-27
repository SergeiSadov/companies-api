package company

import (
	"context"

	"companies-api/internal/entities/api"
	"companies-api/internal/pkg/kafka_writer"
	"companies-api/internal/repositories/company"
	"companies-api/internal/services/company/adapter"
)

type IService interface {
	Create(ctx context.Context, req *api.CreateCompanyRequest) (response *api.CreateCompanyResponse, err error)
	Get(ctx context.Context, req *api.GetCompanyRequest) (response *api.GetCompanyResponse, err error)
	List(ctx context.Context, req *api.ListCompanyRequest) (response []api.Company, err error)
	Count(ctx context.Context, req *api.ListCompanyRequest) (count int, err error)
	Update(ctx context.Context, req *api.UpdateCompanyRequest) (response *api.UpdateCompanyResponse, err error)
	Delete(ctx context.Context, req *api.DeleteCompanyRequest) (err error)
}

type service struct {
	companyRepo  company.IRepository
	createWriter kafka_writer.IKafkaWriter
	updateWriter kafka_writer.IKafkaWriter
	deleteWriter kafka_writer.IKafkaWriter
	adapter      adapter.IAdapter
}

func NewService(
	companyRepo company.IRepository,
	createWriter kafka_writer.IKafkaWriter,
	updateWriter kafka_writer.IKafkaWriter,
	deleteWriter kafka_writer.IKafkaWriter,
	adapter adapter.IAdapter,
) *service {
	return &service{
		companyRepo:  companyRepo,
		createWriter: createWriter,
		updateWriter: updateWriter,
		deleteWriter: deleteWriter,
		adapter:      adapter,
	}
}

func (s service) Create(ctx context.Context, req *api.CreateCompanyRequest) (response *api.CreateCompanyResponse, err error) {
	res, err := s.companyRepo.Create(ctx, s.adapter.AdaptCreateReqToRepo(req))
	if err != nil {
		return
	}

	event, err := s.adapter.AdaptCompanyRepoToKafka(res)
	if err != nil {
		return
	}
	if err = s.createWriter.WriteMessages(ctx, event); err != nil {
		return
	}

	return s.adapter.AdaptRepoToCreateResp(res), nil
}

func (s service) Get(ctx context.Context, req *api.GetCompanyRequest) (response *api.GetCompanyResponse, err error) {
	res, err := s.companyRepo.Get(ctx, s.adapter.AdaptGetReqToRepo(req))
	if err != nil {
		return
	}

	return s.adapter.AdaptRepoToGetResp(res), nil
}

func (s service) List(ctx context.Context, req *api.ListCompanyRequest) (response []api.Company, err error) {
	res, err := s.companyRepo.List(ctx, s.adapter.AdaptListReqToRepo(req))
	if err != nil {
		return
	}

	return s.adapter.AdaptRepoToListResp(res), nil
}

func (s service) Count(ctx context.Context, req *api.ListCompanyRequest) (count int, err error) {
	res, err := s.companyRepo.Count(ctx, s.adapter.AdaptListReqToRepo(req))
	if err != nil {
		return
	}

	return res, nil
}

func (s service) Update(ctx context.Context, req *api.UpdateCompanyRequest) (response *api.UpdateCompanyResponse, err error) {
	res, err := s.companyRepo.Update(ctx, s.adapter.AdaptUpdateReqToRepo(req))
	if err != nil {
		return
	}

	event, err := s.adapter.AdaptCompanyRepoToKafka(res)
	if err != nil {
		return
	}
	if err = s.updateWriter.WriteMessages(ctx, event); err != nil {
		return
	}

	return s.adapter.AdaptRepoToUpdateResp(res), nil
}

func (s service) Delete(ctx context.Context, req *api.DeleteCompanyRequest) (err error) {
	err = s.companyRepo.Delete(ctx, s.adapter.AdaptDeleteReqToRepo(req))
	if err != nil {
		return
	}

	event, err := s.adapter.AdaptIDEventToKafka(req.ID)
	if err != nil {
		return
	}
	if err = s.deleteWriter.WriteMessages(ctx, event); err != nil {
		return
	}

	return nil
}
