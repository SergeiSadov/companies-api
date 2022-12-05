package postgresql

import (
	"companies-api/internal/entities/repository"
	"context"
	"reflect"
	"testing"
)

func TestCompany_Create(t *testing.T) {
	type args struct {
		ctx context.Context
		req *repository.Company
	}
	tests := []struct {
		name         string
		c            *Company
		args         args
		wantResponse *repository.Company
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.Create(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Company.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Company.Create() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestCompany_Get(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name         string
		c            *Company
		args         args
		wantResponse *repository.Company
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.Get(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Company.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Company.Get() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestCompany_Count(t *testing.T) {
	type args struct {
		ctx context.Context
		req *repository.ListCompanyParams
	}
	tests := []struct {
		name         string
		c            *Company
		args         args
		wantResponse int
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.Count(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Company.Count() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResponse != tt.wantResponse {
				t.Errorf("Company.Count() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestCompany_List(t *testing.T) {
	type args struct {
		ctx context.Context
		req *repository.ListCompanyParams
	}
	tests := []struct {
		name         string
		c            *Company
		args         args
		wantResponse []repository.Company
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.List(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Company.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Company.List() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestCompany_Update(t *testing.T) {
	type args struct {
		ctx context.Context
		req *repository.Company
	}
	tests := []struct {
		name         string
		c            *Company
		args         args
		wantResponse *repository.Company
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResponse, err := tt.c.Update(tt.args.ctx, tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Company.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResponse, tt.wantResponse) {
				t.Errorf("Company.Update() = %v, want %v", gotResponse, tt.wantResponse)
			}
		})
	}
}

func TestCompany_Delete(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int
	}
	tests := []struct {
		name    string
		c       *Company
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Delete(tt.args.ctx, tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("Company.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
