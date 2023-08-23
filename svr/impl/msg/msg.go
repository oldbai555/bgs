package msg

import (
	"github.com/oldbai555/bgs/svr/impl/processor"
)

// 消息注册使用

var Processor *processor.Processor

func init() {
	Processor = processor.NewProcessor()
	for i := range msgList {
		Processor.Register(msgList[i])
	}
}

var msgList []interface{}
