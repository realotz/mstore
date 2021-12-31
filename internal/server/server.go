package server

import (
	"encoding/json"
	"github.com/go-kratos/kratos/v2/encoding"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/google/wire"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"net/http"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, NewCronServer)

var (
	// MarshalOptions is a configurable JSON format marshaler.
	MarshalOptions = protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
	// UnmarshalOptions is a configurable JSON format parser.
	UnmarshalOptions = protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}
)

func init() {
	encoding.RegisterCodec(codec{})
}

// codec json解析器
type codec struct{}

type Success struct {
	Code    uint32      `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (codec) Marshal(v interface{}) ([]byte, error) {
	if m, ok := v.(*errors.Error); ok {
		return MarshalOptions.Marshal(m)
	}
	if m, ok := v.(proto.Message); ok {
		pbBytes, err := MarshalOptions.Marshal(m)
		if err != nil {
			return nil, err
		}
		return jsoniter.Marshal(Success{
			http.StatusOK,
			json.RawMessage(pbBytes),
			"ok",
		})
	}
	return jsoniter.Marshal(v)
}

func (codec) Unmarshal(data []byte, v interface{}) error {
	if m, ok := v.(proto.Message); ok {
		return UnmarshalOptions.Unmarshal(data, m)
	}
	return jsoniter.Unmarshal(data, v)
}

func (codec) Name() string {
	return "json"
}
