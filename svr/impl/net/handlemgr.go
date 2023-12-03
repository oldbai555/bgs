package net

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/lbtool/log"
)

type HandlePbFunc func(a gate.Agent, msg *pb.Message) error

var (
	pbFuncVec = make(map[uint32]map[uint64]HandlePbFunc)
)

// 注册
func regPbHandleFunc(serverType pb.ServerType) func(high, low uint32, fn HandlePbFunc) {
	return func(high, low uint32, fn HandlePbFunc) {
		cmd := tool.Make64(low, high)
		st := uint32(serverType)
		mgr, ok := pbFuncVec[st]
		if !ok {
			pbFuncVec[st] = make(map[uint64]HandlePbFunc)
			mgr = pbFuncVec[st]
		}
		_, ok = mgr[cmd]
		if ok {
			panic(fmt.Sprintf("already registered handle , low is %d , high is %d", low, high))
		}
		mgr[cmd] = fn
	}
}

// 获取
func getPbHandleFunc(serverType pb.ServerType) func(cmd uint64) (HandlePbFunc, bool) {
	return func(cmd uint64) (HandlePbFunc, bool) {
		st := uint32(serverType)
		mgr, ok := pbFuncVec[st]
		if !ok {
			log.Warnf("not found handle mgr")
			return nil, false
		}
		fn, ok := mgr[cmd]
		if !ok {
			log.Warnf("not found handle")
			return nil, false
		}
		return fn, true
	}
}

// 发送
func sendPbMsg(serverType pb.ServerType) func(agent gate.Agent, high, low uint32, msg proto.Message) error {
	return func(agent gate.Agent, high, low uint32, msg proto.Message) error {
		st := uint32(serverType)
		buf, err := tool.Marshal(msg)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}

		m := &pb.Message{
			ProtoH: high,
			ProtoL: low,
			Data:   string(buf),
		}

		agent.WriteMsg(&pb.Event{
			ServerType: st,
			Message:    m,
		})
		return nil
	}
}

var (
	RegGameSvrPbHandle = regPbHandleFunc(pb.ServerType_ServerTypeGameSvr)
	GetGameSvrPbHandle = getPbHandleFunc(pb.ServerType_ServerTypeGameSvr)

	RegAuthPbHandle = regPbHandleFunc(pb.ServerType_ServerTypeAuthSvr)
	GetAuthPbHandle = getPbHandleFunc(pb.ServerType_ServerTypeAuthSvr)
	SendAuthPbMsg   = sendPbMsg(pb.ServerType_ServerTypeAuthSvr)

	RegLocalSceneSvrPbHandle = regPbHandleFunc(pb.ServerType_ServerTypeLocalSceneSvr)
	GetLFPbHandle            = getPbHandleFunc(pb.ServerType_ServerTypeLocalSceneSvr)
)
