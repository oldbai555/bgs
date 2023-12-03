package base

import (
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/lbtool/log"
)

func SendProto(agent gate.Agent, h, l uint32, msg proto.Message) error {
	bytes, err := tool.Marshal(msg)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	agent.WriteMsg(&pb.Message{
		ProtoH: h,
		ProtoL: l,
		Data:   string(bytes),
	})
	return nil
}
