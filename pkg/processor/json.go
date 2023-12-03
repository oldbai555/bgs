package processor

import (
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/chanrpc"
	"github.com/name5566/leaf/log"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
)

type Processor struct {
	msgInfo map[string]*MsgInfo
}

type MsgInfo struct {
	msgType    string
	msgRouter  *chanrpc.Server
	msgHandler MsgHandler
}

type MsgHandler func([]interface{})

func NewProcessor() *Processor {
	p := new(Processor)
	p.msgInfo = make(map[string]*MsgInfo)
	return p
}

// Register It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) Register(msg fmt.Stringer) string {
	msgType := msg.String()
	if msgType == "" {
		log.Fatal("unnamed processor message")
	}
	if _, ok := p.msgInfo[msgType]; ok {
		log.Fatal("message %v is already registered", msgType)
	}

	i := new(MsgInfo)
	i.msgType = msgType
	p.msgInfo[msgType] = i
	return msgType
}

// SetRouter It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetRouter(msg fmt.Stringer, msgRouter *chanrpc.Server) {
	msgType := msg.String()
	i, ok := p.msgInfo[msgType]
	if !ok {
		log.Fatal("message %v not registered", msgType)
	}

	i.msgRouter = msgRouter
}

// SetHandler It's dangerous to call the method on routing or marshaling (unmarshaling)
func (p *Processor) SetHandler(msg fmt.Stringer, msgHandler MsgHandler) {
	msgType := msg.String()
	i, ok := p.msgInfo[msgType]
	if !ok {
		log.Fatal("message %v not registered", msgType)
	}

	i.msgHandler = msgHandler
}

// Route goroutine safe
func (p *Processor) Route(msg interface{}, userData interface{}) error {
	event, ok := msg.(*pb.Event)
	if !ok {
		return errors.New("invalid processor data")
	}

	if event.ServerType == 0 {
		return errors.New("server type not supported")
	}

	msgType := pb.ServerType_name[int32(event.ServerType)]
	i, ok := p.msgInfo[msgType]
	if !ok {
		return fmt.Errorf("message %v not registered", msgType)
	}

	if i.msgHandler != nil {
		i.msgHandler([]interface{}{event.Message, userData})
	}

	if i.msgRouter != nil {
		i.msgRouter.Go(msgType, event.Message, userData)
	}
	return nil
}

// Unmarshal goroutine safe
func (p *Processor) Unmarshal(data []byte) (interface{}, error) {
	var m pb.Event
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
