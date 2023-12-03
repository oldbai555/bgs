package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/svr/impl/conf"
	"github.com/oldbai555/bgs/svr/impl/moude/auth"
	"github.com/oldbai555/bgs/svr/impl/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		CertFile:        conf.Server.CertFile,
		KeyFile:         conf.Server.KeyFile,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.GetSingeProcess(), // 指向全局的processor
		AgentChanRPC:    auth.ChanRPC,          // 指向认证模块的AgentChanRPC
	}
}
