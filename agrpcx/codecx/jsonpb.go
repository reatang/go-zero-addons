package codecx

import (
	"bytes"
	"encoding/json"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc/encoding"
)

const Json = "json"

func RegisterCodecJson() {
	encoding.RegisterCodec(JSONCodec{
		Marshaler: jsonpb.Marshaler{
			EmitDefaults: true,
			OrigName:     true,
		},
		Unmarshaler: jsonpb.Unmarshaler{
			AllowUnknownFields: false,
			AnyResolver:        nil,
		},
	})
}

type JSONCodec struct {
	jsonpb.Marshaler
	jsonpb.Unmarshaler
}

func (_ JSONCodec) Name() string {
	return Json
}

func (j JSONCodec) Marshal(v interface{}) (out []byte, err error) {
	if pm, ok := v.(proto.Message); ok {
		b := new(bytes.Buffer)
		err := j.Marshaler.Marshal(b, pm)
		if err != nil {
			return nil, err
		}
		return b.Bytes(), nil
	}
	return json.Marshal(v)
}

func (j JSONCodec) Unmarshal(data []byte, v interface{}) (err error) {
	if pm, ok := v.(proto.Message); ok {
		b := bytes.NewBuffer(data)
		return j.Unmarshaler.Unmarshal(b, pm)
	}
	return json.Unmarshal(data, v)
}
