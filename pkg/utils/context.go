package utils

import (
	"context"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/realotz/mstore/pkg/myctx"
	"strings"
)

func GetClientIp(ctx context.Context) string {
	tr, ok := transport.FromServerContext(ctx)
	ct, _ := myctx.FormHttpContext(ctx)
	if ct != nil {
		tr.RequestHeader().Set("Remote-Addr", ct.Request.RemoteAddr)
	}
	if ok {
		clientIP := tr.RequestHeader().Get("X-Forwarded-For")
		clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
		if clientIP == "" {
			clientIP = strings.TrimSpace(tr.RequestHeader().Get("X-Real-Ip"))
		}
		if clientIP != "" {
			return clientIP
		}
		if addr := tr.RequestHeader().Get("X-Appengine-Remote-Addr"); addr != "" {
			return addr
		}
		if addr := tr.RequestHeader().Get("Remote-Addr"); addr != "" {
			return addr
		}
	}
	return "127.0.0.1"
}
