package error_encoders

import (
	"context"
	"net/http"

	"companies-api/internal/pkg/error_adapters/http_adapter"

	"go.uber.org/zap"
)

func NewErrorEncoder(errorAdapter http_adapter.IErrorAdapter, logger *zap.Logger) func(_ context.Context, err error, w http.ResponseWriter) {
	return func(_ context.Context, err error, w http.ResponseWriter) {
		code := errorAdapter.AdaptToHttpCode(err)
		if code != http.StatusBadRequest {
			logger.Error(err.Error())
		}

		w.WriteHeader(code)
	}
}
