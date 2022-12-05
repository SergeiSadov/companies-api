package validators

import (
	"companies-api/internal/entities/api"
	"testing"
)

func TestValidators_ValidateCreateRequest(t *testing.T) {
	type args struct {
		req *api.CreateCompanyRequest
	}
	tests := []struct {
		name    string
		v       *Validators
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.ValidateCreateRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Validators.ValidateCreateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidators_ValidateGetRequest(t *testing.T) {
	type args struct {
		req *api.GetCompanyRequest
	}
	tests := []struct {
		name    string
		v       *Validators
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.ValidateGetRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Validators.ValidateGetRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidators_ValidateUpdateRequest(t *testing.T) {
	type args struct {
		req *api.UpdateCompanyRequest
	}
	tests := []struct {
		name    string
		v       *Validators
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.ValidateUpdateRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Validators.ValidateUpdateRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidators_ValidateDeleteRequest(t *testing.T) {
	type args struct {
		req *api.DeleteCompanyRequest
	}
	tests := []struct {
		name    string
		v       *Validators
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.ValidateDeleteRequest(tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("Validators.ValidateDeleteRequest() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
