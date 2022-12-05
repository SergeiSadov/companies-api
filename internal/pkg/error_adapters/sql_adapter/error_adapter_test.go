package sql_adapter

import (
	"companies-api/internal/pkg/errors"
	"fmt"
	"testing"

	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestErrorAdapter_AdaptSqlErr(t *testing.T) {
	adapter := New(errors.PreparedPostrgesErrorsMap)

	type args struct {
		err error
	}
	tests := []struct {
		name    string
		a       *ErrorAdapter
		err     error
		wantErr error
	}{
		{
			"UniqueViolationErr",
			adapter,
			&pq.Error{
				Code: errors.UniqueViolationErr,
			},
			errors.ErrAlreadyExist,
		},
		{
			"ErrNotFound",
			adapter,
			&pq.Error{
				Code: errors.NotFoundErr,
			},
			errors.ErrNotFound,
		},
		{
			"Return unadapted err",
			adapter,
			fmt.Errorf("error"),
			fmt.Errorf("error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.wantErr, tt.a.AdaptSqlErr(tt.err))
		})
	}
}
