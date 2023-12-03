package internal

import (
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/lbtool/log"
)

func regHandle() {
	registerChanRPC(2, 1, c2sLogin)
}

func c2sLogin(agent gate.Agent, msg *pb.Message) error {
	var req pb.C2S_2_1
	err := tool.Unmarshal([]byte(msg.Data), &req)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	agent.WriteMsg(&pb.S2C_2_1{
		Username: req.Username,
		Password: req.Password,
	})
	return nil
}
