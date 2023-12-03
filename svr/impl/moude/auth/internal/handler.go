package internal

import (
	"context"
	"fmt"
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/pkg/gorm"
	"github.com/oldbai555/bgs/pkg/tool"
	"github.com/oldbai555/bgs/proto/pb"
	"github.com/oldbai555/bgs/svr/impl/base"
	"github.com/oldbai555/bgs/svr/impl/net"
	"github.com/oldbai555/lbtool/log"
	"github.com/oldbai555/lbtool/utils"
)

func handleMsg() {
	skeleton.RegisterChanRPC(pb.ServerType_ServerTypeAuthSvr.String(), handle)
}

func init() {
	handleMsg()
	net.RegAuthPbHandle(2, 1, c2sLogin)
}

func handle(args []interface{}) interface{} {
	message, ok := args[0].(*pb.Message)
	if !ok {
		return fmt.Errorf("event parse failed")
	}

	a, ok := args[1].(gate.Agent)
	if !ok {
		return fmt.Errorf("agent parse failed")
	}

	cmd := tool.Make64(message.ProtoL, message.ProtoH)

	handle, ok := net.GetAuthPbHandle(cmd)
	if !ok {
		log.Warnf("not found handler %v", cmd)
		a.WriteMsg(base.ErrNotFoundHandle)
		return nil
	}

	err := handle(a, message)
	if err != nil {
		log.Errorf("err:%v", err)
		a.WriteMsg(base.RegClientErr(uint32(pb.ErrCode_ErrSystemError), err.Error()))
		return err
	}

	return nil
}

func c2sLogin(a gate.Agent, msg *pb.Message) error {
	var req pb.C2S_2_1
	err := tool.Unmarshal([]byte(msg.Data), &req)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	ctx := context.Background()
	err = gorm.NewScope(ctx, &pb.ModelAccount{}).Where(map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	}).First(ctx, &pb.ModelAccount{})
	if err != nil && !gorm.IsNotFoundErr(err) {
		log.Errorf("err:%v", err)
		return err
	}

	if gorm.IsNotFoundErr(err) {
		_, err = gorm.NewScope(ctx, &pb.ModelAccount{}).Create(ctx, &pb.ModelAccount{
			Username: req.Username,
			Password: req.Password,
		})
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
	}

	_, err = gorm.NewScope(ctx, &pb.ModelAccount{}).Where(map[string]interface{}{
		"username": req.Username,
		"password": req.Password,
	}).Update(ctx, map[string]interface{}{
		"last_login_ip": a.RemoteAddr().String(),
		"last_login_at": utils.TimeNow(),
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	err = net.SendAuthPbMsg(a, 2, 1, &pb.S2C_2_1{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}
	return nil
}
