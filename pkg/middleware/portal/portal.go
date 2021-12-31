package portal

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"git.hxecloud.com/cloudwonder-portal/rms-server/pkg/utils"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	http2 "net/http"
	"time"
)

type OptionFunc func(op *option)

type option struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
}

type ErrorReply struct {
	Errors []struct {
		Field string `json:"field"`
		Desc  string `json:"desc"`
	} `json:"errors"`
	Msg string `json:"msg"`
}

func ErrorDecoder(ctx context.Context, res *http2.Response) error {
	if res.StatusCode >= 200 && res.StatusCode <= 299 {
		return nil
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err == nil {
		var resp = ErrorReply{}
		_ = jsoniter.Unmarshal(data, &resp)
		return errors.Errorf(res.StatusCode, resp.Msg, utils.JsonString(resp.Errors))
	}
	return errors.Errorf(res.StatusCode, errors.UnknownReason, "")
}

func WithConfig(appId, appSecret string) OptionFunc {
	return func(op *option) {
		op.AppId = appId
		op.AppSecret = appSecret
	}
}

func genSign(timestamp int64, appID, appSecret string) string {
	sigStr := fmt.Sprintf("%v&%v&%v", appID, timestamp, appSecret)
	sigStrSum := sha256.Sum256([]byte(sigStr))
	return base64.RawURLEncoding.EncodeToString(sigStrSum[:])
}

func Client(ops ...OptionFunc) middleware.Middleware {
	op := &option{}
	for _, v := range ops {
		v(op)
	}
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if info, ok := transport.FromClientContext(ctx); ok {
				timestamp := time.Now().Unix()
				info.RequestHeader().Set("x-appid", op.AppId)
				info.RequestHeader().Set("x-timestamp", fmt.Sprintf("%v", timestamp))
				info.RequestHeader().Set("x-authorization", genSign(timestamp, op.AppId, op.AppSecret))
			}
			return handler(ctx, req)
		}
	}
}
