package ftp

import (
	"strings"

	"github.com/r00tu53r/protos/statemachine"
)

var stateMachines map[string]*statemachine.StateMachine

func init() {
	stateMachines = make(map[string]*statemachine.StateMachine)
	lsm := newLoginStateMachine()
	stateMachines[User] = lsm
	stateMachines[Password] = lsm
	stateMachines[Account] = lsm
}

func getStateMachine(cmd string) (*statemachine.StateMachine, error) {
	_, ok := Commands[strings.ToUpper(cmd)]
	if !ok {
		return nil, ErrStateMachineNotFound
	}
	return stateMachines[strings.ToUpper(cmd)], nil
}

func toks(cmds []string) [][]byte {
	toks := make([][]byte, len(cmds))
	for i := range cmds {
		toks[i] = append(toks[i], []byte(cmds[i])...)
	}
	return toks
}

func newLoginStateMachine() *statemachine.StateMachine {
	sm := statemachine.New()
	sm.AddState(
		StateBegin, nil,
		StateWaitForReply, nil,
		toks([]string{User}),
	)
	sm.AddState(
		StateWaitForReply, nil,
		StateNeedPassword, nil,
		toks([]string{"331"}),
	)
	sm.AddState(
		StateWaitForReply, nil,
		StateNeedPassword, nil,
		toks([]string{"331"}),
	)
	sm.AddState(
		StateNeedPassword, nil,
		StateWaitForReply, nil,
		toks([]string{Password}),
	)
	sm.AddState(
		StateWaitForReply, nil,
		StateSuccess, nil,
		toks([]string{"230"}),
	)
	sm.AddState(
		StateWaitForReply, nil,
		StateFailed, nil,
		toks([]string{"530"}),
	)
	return sm
}
