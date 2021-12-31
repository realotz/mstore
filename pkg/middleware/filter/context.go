package filter

import (
	"github.com/realotz/mstore/pkg/myctx"
	"net/http"
)

func HttpContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := myctx.NewHttpServerContext(r.Context(), &myctx.Context{
			Writer:  w,
			Request: r,
		})
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
