package net

import (
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/lbtool/log"
)

type HandleProtoFunc func(agent gate.Agent, msg *pb.Message) error

func PackingFunction(handle HandleProtoFunc) func([]interface{}) interface{} {
	return func(args []interface{}) interface{} {
		agent, ok := args[0].(gate.Agent)
		if !ok {
			panic("agent parse failed")
		}

		message, ok := args[1].(*pb.Message)
		if !ok {
			panic("event parse failed")
		}

		err := handle(agent, message)
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}

		return nil
	}
}
