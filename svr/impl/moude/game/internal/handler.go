package internal

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/bgs/svr/impl/net"
	"github.com/oldbai555/lbtool/log"
)

func handleMsg() {
	skeleton.RegisterChanRPC(pb.ServerType_ServerTypeGameSvr.String(), handle)
}

func init() {
	handleMsg()
}

func handle(args []interface{}) interface{} {
	message, ok := args[0].(*pb.Message)
	if !ok {
		return fmt.Errorf("event parse failed")
	}

	a, ok := args[1].(gate.Agent)
	if !ok {
		return fmt.Errorf("agent parse failed")
	}

	cmd := tool.Make64(message.ProtoL, message.ProtoH)

	handle, ok := net.GetGameSvrPbHandle(cmd)
	if !ok {
		log.Warnf("not found handle %v", cmd)
		return nil
	}

	err := handle(a, message)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}
