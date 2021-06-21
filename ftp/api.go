package ftp

import (
	"strconv"

	"github.com/r00tu53r/protos/ftp/fsm"
)

func Run(message string, currentState *fsm.FSMState) (*fsm.FSMState, error) {
	if message == "" {
		return nil, fsm.ErrInvalidMessage
	}
	reqOp, isReq := isRequest(message)
	respCode, isResp := isResponse(message)
	if !isReq && !isResp {
		return nil, fsm.ErrInvalidMessage
	}
	if currentState == nil {
		currentState = &fsm.FSMState{}
	}
	currentState.IsRequest = isReq
	if isReq {
		currentState.ReqOp = reqOp
	} else {
		currentState.RespCode = strconv.Itoa(respCode)
	}
	sm, err := fsm.GetStateMachine(currentState.ReqOp)
	if err != nil {
		return currentState, err
	}
	newState, err := sm.Transition(currentState, message)
	return newState, err
}
