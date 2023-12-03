package syscall

import (
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/lbtool/log"
)

type HandleSysFunc func(a gate.Agent, msg *pb.Message) error

var (
	sysFuncVec = make(map[uint64]map[uint64]HandleSysFunc)
)

func regSysCallFunc(serverType pb.ServerType) func(cmd uint64, fn HandleSysFunc) {
	return func(cmd uint64, fn HandleSysFunc) {
		st := uint64(serverType)
		mgr, ok := sysFuncVec[st]
		if !ok {
			sysFuncVec[st] = make(map[uint64]HandleSysFunc)
			mgr = sysFuncVec[st]
		}
		_, ok = mgr[cmd]
		if ok {
			panic(fmt.Sprintf("already registered handle , cmd is %d ", cmd))
		}
		mgr[cmd] = fn
		log.Infof("registered handle server type %d , cmd is %d ", cmd)
	}
}

func getSysCallFunc(serverType pb.ServerType) func(cmd uint64) (HandleSysFunc, bool) {
	return func(cmd uint64) (HandleSysFunc, bool) {
		st := uint64(serverType)
		mgr, ok := sysFuncVec[st]
		if !ok {
			log.Warnf("not found handle , server type is %d ", serverType)
			return nil, false
		}
		handle, ok := mgr[cmd]
		if !ok {
			log.Warnf("not found handle , cmd is %d ", cmd)
			return nil, false
		}
		return handle, true
	}
}

var (
	GetGSSysCallFunc = getSysCallFunc(pb.ServerType_ServerTypeGameSvr)
	RegGSSysCallFunc = regSysCallFunc(pb.ServerType_ServerTypeGameSvr)
)
