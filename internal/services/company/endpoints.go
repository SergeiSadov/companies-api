package company

import (
	"context"
	"math"

	"companies-api/internal/entities/api"
	"companies-api/internal/services/company/validators"

	"github.com/go-kit/kit/endpoint"
)

func makeListCompaniesEndpoint(svc IService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.ListCompanyRequest)

		data, err := svc.List(ctx, req)
		if err != nil {
			return nil, err
		}

		count, err := svc.Count(ctx, req)
		if err != nil {
			return nil, err
		}

		return &api.ListCompanyResponse{
			Data: data,
			Meta: api.Meta{
				Pages: int(math.Ceil(float64(count) / float64(req.Size))),
				Total: count,
			},
		}, nil
	}
}

func makeCreateCompanyEndpoint(svc IService, validator validators.IValidator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.CreateCompanyRequest)

		if err = validator.ValidateCreateRequest(req); err != nil {
			return nil, err
		}

		data, err := svc.Create(ctx, req)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
}

func makeGetCompanyEndpoint(svc IService, validator validators.IValidator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.GetCompanyRequest)

		if err = validator.ValidateGetRequest(req); err != nil {
			return nil, err
		}

		data, err := svc.Get(ctx, req)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
}

func makeUpdateCompanyEndpoint(svc IService, validator validators.IValidator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.UpdateCompanyRequest)

		if err = validator.ValidateUpdateRequest(req); err != nil {
			return nil, err
		}

		data, err := svc.Update(ctx, req)
		if err != nil {
			return nil, err
		}

		return data, nil
	}
}

func makeDeleteCompanyEndpoint(svc IService, validator validators.IValidator) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*api.DeleteCompanyRequest)

		if err = validator.ValidateDeleteRequest(req); err != nil {
			return nil, err
		}

		err = svc.Delete(ctx, req)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
}
