package impl

import (
	"github.com/name5566/leaf"
	"github.com/oldbai555/bgs/svr/impl/conf"
	"github.com/oldbai555/bgs/svr/impl/moude"
	"github.com/oldbai555/bgs/svr/impl/msg"
	"github.com/oldbai555/lbtool/log"

	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	conf.LoadServerConfig("")
	log.SetModuleName("bgs")

	//gorm.RegisterModel(
	//	&pb.ModelAccount{},
	//	&pb.ModelService{},
	//	&pb.ModelPlatform{},
	//	&pb.ModelActor{},
	//)
	//err := gorm.InitGorm("")
	//if err != nil {
	//	log.Errorf("err:%v", err)
	//	return err
	//}

	msg.InitMsgProcess()

	leaf.Run(moude.Modules()...)
	return nil
}
