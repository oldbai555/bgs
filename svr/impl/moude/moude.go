package moude

import (
	"github.com/name5566/leaf/module"
	"github.com/oldbai555/bgs/svr/impl/moude/auth"
	"github.com/oldbai555/bgs/svr/impl/moude/game"
	"github.com/oldbai555/bgs/svr/impl/moude/gate"
)

func Modules() []module.Module {
	return []module.Module{
		auth.Module,
		gate.Module,
		game.Module,
	}
}
