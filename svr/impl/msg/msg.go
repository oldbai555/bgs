package msg

import (
	"github.com/oldbai555/bgs/pkg/processor"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/bgs/svr/impl/moude/auth"
	"github.com/oldbai555/bgs/svr/impl/moude/game"
	"sync"
)

// 消息注册使用

var process *processor.Processor
var once sync.Once

func InitMsgProcess() {
	singeProcess := GetSingeProcess()
	singeProcess.Register(pb.ServerType_ServerTypeAuthSvr)
	singeProcess.SetRouter(pb.ServerType_ServerTypeAuthSvr, auth.ChanRPC)

	singeProcess.Register(pb.ServerType_ServerTypeGameSvr)
	singeProcess.SetRouter(pb.ServerType_ServerTypeGameSvr, game.ChanRPC)
}

func GetSingeProcess() *processor.Processor {
	if process == nil {
		once.Do(func() {
			process = processor.NewProcessor()
		})
	}
	return process
}
