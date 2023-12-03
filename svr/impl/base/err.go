package base

import (
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
)

func regErr(svrType pb.ServerType, code uint32, desc string) func() *pb.Event {
	return func() *pb.Event {
		bytes, _ := tool.Marshal(&pb.S2C_1_1{
			ErrCode: code,
			Msg:     desc,
		})
		return &pb.Event{
			ServerType: uint32(svrType),
			Message: &pb.Message{
				ProtoH: 1,
				ProtoL: 1,
				Data:   string(bytes),
			},
		}
	}
}

func regClientErr(code uint32, desc string) func() *pb.Event {
	return regErr(pb.ServerType_ServerTypeClient, code, desc)
}

var (
	ErrNotFoundHandle = regClientErr(uint32(pb.ErrCode_ErrNotFoundHandle), "not found handle")
	ErrNotFoundConf   = regClientErr(uint32(pb.ErrCode_ErrNotFoundConf), "not found conf")
	ErrInvalidArgs    = regClientErr(uint32(pb.ErrCode_ErrInvalidArgs), "invalid args")
	RegErr            = regErr
	RegClientErr      = regClientErr
)
