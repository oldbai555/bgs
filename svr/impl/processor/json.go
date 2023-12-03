package processor

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/bgs/svr/impl/engine"
)

type Processor struct {
	msgInfo     map[string]*MsgInfo
	cmd2SvrType map[uint64]pb.ServerType
}

type MsgInfo struct {
	msgType   string
	msgRouter *chanrpc.Server
}

func NewProcessor() *Processor {
	p := new(Processor)
	p.msgInfo = make(map[string]*MsgInfo)
	p.cmd2SvrType = make(map[uint64]pb.ServerType)
	engine.RegCmd2SrvTyp = p.regCmd2SvrType
	return p
}

func (p *Processor) RegRouter(msg fmt.Stringer, msgRouter *chanrpc.Server) {
	msgType := msg.String()
	if msgType == "" {
		panic("unnamed processor message")
	}
	if _, ok := p.msgInfo[msgType]; ok {
		panic(fmt.Sprintf("message %v is already registered", msgType))
	}

	i := new(MsgInfo)
	i.msgType = msgType
	i.msgRouter = msgRouter
	p.msgInfo[msgType] = i
}

func (p *Processor) regCmd2SvrType(cmdId uint64, serverType pb.ServerType) {
	if _, ok := p.cmd2SvrType[cmdId]; ok {
		panic(fmt.Sprintf("cmdId %v is already registered", cmdId))
	}
	p.cmd2SvrType[cmdId] = serverType
}

// Route goroutine safe
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	message, ok := msg.(*pb.Message)
	if !ok {
		return fmt.Errorf("invalid processor data")
	}

	cmdId := tool.Make64(message.ProtoL, message.ProtoH)
	msgType, ok := p.cmd2SvrType[cmdId]
	if !ok {
		return fmt.Errorf("protoH %d , protoL %d not registered", message.ProtoH, message.ProtoL)
	}

	i, ok := p.msgInfo[msgType.String()]
	if !ok {
		return fmt.Errorf("message %v not registered", msgType)
	}

	if i.msgRouter != nil {
		i.msgRouter.Go(cmdId, userData, message)
	}
	return nil
}

// Unmarshal goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	var m pb.Message
	err := tool.Unmarshal(data, &m)
	if err != nil {
		return nil, err
	}
	return &m, nil
}

// Marshal goroutine safe
func (p *Processor) Marshal(msg interface{}) ([][]byte, error) {
	data, err := tool.Marshal(msg.(proto.Message))
	return [][]byte{data}, err
}
