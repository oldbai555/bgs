package impl

import (
	"github.com/name5566/leaf"
	"github.com/oldbai555/bgs/svr/impl/conf"
	"github.com/oldbai555/bgs/svr/impl/moude"

	"github.com/urfave/cli/v2"
)

func Run(ctx *cli.Context) error {
	conf.LoadServerConfig("")

	leaf.Run(moude.Modules()...)
	return nil
}
