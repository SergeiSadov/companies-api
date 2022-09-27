//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
package middlewares

import (
	"net/http"
)

type IMiddleware interface {
	HTTPMiddleware(h http.Handler) http.Handler
}
