package tool

import (
	"bytes"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/oldbai555/lbtool/log"
)

func Unmarshal(data []byte, val proto.Message) error {
	unmarshaler := &jsonpb.Unmarshaler{AllowUnknownFields: true}
	err := unmarshaler.Unmarshal(bytes.NewReader(data), val)
	if err != nil {
		return err
	}
	return nil
}

func Marshal(val proto.Message) ([]byte, error) {
	marshal := &jsonpb.Marshaler{}
	var buf []byte
	w := bytes.NewBuffer(buf)
	err := marshal.Marshal(w, val)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}
	return w.Bytes(), nil
}
