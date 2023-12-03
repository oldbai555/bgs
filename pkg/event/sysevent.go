package event

import (
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/pkg/routine"
)

type SysECallBack func(args ...interface{}) error

var (
	sysEFuncVec = make(map[uint32][]SysECallBack)
)

func TriggerSysEvent(id uint32, args ...interface{}) {
	for _, fn := range sysEFuncVec[id] {
		routine.GoV2(func() error {
			err := fn(args...)
			if err != nil {
				log.Errorf("err:%v", err)
				return err
			}
			return nil
		})
	}
}

func regSysEFunc() func(id uint32, fn SysECallBack) {
	return func(id uint32, fn SysECallBack) {
		if id == 0 {
			panic("id is zero")
		}
		if fn == nil {
			panic("fn is nil")
		}
		sysEFuncVec[id] = append(sysEFuncVec[id], fn)
	}
}

var (
	RegSysEvent = regSysEFunc()
)
