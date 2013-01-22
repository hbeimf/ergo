package node

import (
	erl "github.com/goerlang/etf/types"
)

type globalNameServer struct {
	gsi *GenServerImpl
}

func (gns *globalNameServer) Behaviour() (behaviour Behaviour, options map[string]interface{}) {
	gns.gsi = &GenServerImpl{}
	return gns.gsi, gns.gsi.Options()
}

func (gns *globalNameServer) Init(args ...interface{}) {
	nLog("GLOBAL_NAME_SERVER: Init: %#v", args)
	gns.gsi.node.Register(erl.Atom("global_name_server"), gns.gsi.self)
}

func (gns *globalNameServer) HandleCast(message *erl.Term) {
	nLog("GLOBAL_NAME_SERVER: HandleCast: %#v", *message)
}

func (gns *globalNameServer) HandleCall(message *erl.Term, from *erl.Tuple) (reply *erl.Term) {
	nLog("GLOBAL_NAME_SERVER: HandleCall: %#v, From: %#v", *message, *from)
	replyTerm := erl.Term(erl.Atom("reply"))
	reply = &replyTerm
	return
}

func (gns *globalNameServer) HandleInfo(message *erl.Term) {
	nLog("GLOBAL_NAME_SERVER: HandleInfo: %#v", *message)
}

func (gns *globalNameServer) Terminate(reason interface{}) {
	nLog("GLOBAL_NAME_SERVER: Terminate: %#v", reason.(int))
}
