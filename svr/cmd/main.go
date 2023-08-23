package main

import (
	"github.com/oldbai555/bgs/svr/impl"
	"github.com/oldbai555/lbtool/log"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "bgs"
	app.Version = "v0.0.1"
	app.Description = "bgs ..."
	app.Action = impl.Run
	err := app.Run(os.Args)
	if err != nil {
		log.Errorf("err:%v", err)
		return
	}
}
