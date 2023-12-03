package processor

import (
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/bgs/svr/impl/moude/game"
	"sync"
)

var process *Processor
var once sync.Once

func Init() {
	singeProcess := GetOne()
	singeProcess.RegRouter(pb.ServerType_ServerTypeGameSvr, game.ChanRPC)
}

func GetOne() *Processor {
	if process == nil {
		once.Do(func() {
			process = NewProcessor()
		})
	}
	return process
}
