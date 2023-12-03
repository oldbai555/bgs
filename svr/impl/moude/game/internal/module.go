package internal

import (
	"github.com/name5566/leaf/module"
	"github.com/oldbai555/bgs/pkg/net"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/bgs/svr/impl/base"
	"github.com/oldbai555/bgs/svr/impl/engine"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	regHandle()
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}

func registerChanRPC(h, l uint32, protoFunc net.HandleProtoFunc) {
	cmdId := tool.Make64(l, h)
	engine.RegCmd2SrvTyp(cmdId, pb.ServerType_ServerTypeGameSvr)
	skeleton.RegisterChanRPC(cmdId, net.PackingFunction(protoFunc))
}
