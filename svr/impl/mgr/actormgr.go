package mgr

import (
	"github.com/name5566/leaf/gate"
	"github.com/oldbai555/bgs/pkg/tool"
)

type ActorAgentMgr struct {
	M   map[uint64]gate.Agent
	IpM map[int64]uint64
}

func (m *ActorAgentMgr) Reg(actorId uint64, agent gate.Agent) {
	m.M[actorId] = agent
	ip := tool.Ip2int(agent.RemoteAddr().String())
	m.IpM[ip] = actorId
}

func (m *ActorAgentMgr) GetAgent(actorId uint64) (gate.Agent, bool) {
	agent, ok := m.M[actorId]
	return agent, ok
}

func (m *ActorAgentMgr) GetActorId(agent gate.Agent) (uint64, bool) {
	ip := tool.Ip2int(agent.RemoteAddr().String())
	actorId, ok := m.IpM[ip]
	return actorId, ok
}

func (m *ActorAgentMgr) Each(fn func(actorId uint64, agent gate.Agent)) {
	for actorId, agent := range m.M {
		fn(actorId, agent)
	}
}
