package moude

import (
	"github.com/name5566/leaf/module"
	"github.com/oldbai555/bgs/svr/impl/moude/game"
	"github.com/oldbai555/bgs/svr/impl/moude/gate"
	"github.com/oldbai555/bgs/svr/impl/moude/login"
)

func Modules() []module.Module {
	return []module.Module{
		game.Module,
		gate.Module,
		login.Module,
	}
}
