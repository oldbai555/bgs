package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/lbtool/log"
)

const (
	AgentByNew   = "NewAgent"
	AgentByClose = "CloseAgent"
)

func init() {
	skeleton.RegisterChanRPC(AgentByNew, rpcNewAgent)
	skeleton.RegisterChanRPC(AgentByClose, rpcCloseAgent)
}

// agent 被创建时
func rpcNewAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Infof("new agent local addr %s, remote addr %s", a.LocalAddr(), a.RemoteAddr())
}

// agent 被关闭时
func rpcCloseAgent(args []interface{}) {
	a := args[0].(gate.Agent)
	log.Infof("close agent local addr %s, remote addr %s", a.LocalAddr(), a.RemoteAddr())
}
