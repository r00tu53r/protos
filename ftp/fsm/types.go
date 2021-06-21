package fsm

import "errors"

type State int

// RFC959 uses these basic states to represent
// most of the operations.
const (
	StateBegin State = iota
	StateWaitForReply
	StateNeedPassword
	StateNeedAccount
	StateSuccess
	StateError
	StateFailed
)

// fsmState represents the finite state machine
// transitions from a given current state.
type fsmState struct {
	tokens    []string
	next      State
	isRequest bool
}

type FSMState struct {
	Current   State
	Err       error
	IsRequest bool
	ReqOp     string
	RespCode  string
}

// Holds states (tokens and next state) for the
// current state
type stateMachine struct {
	// TODO []State to map[FSMState][]string
	// so we can search by the nextState?
	states map[State][]fsmState
}

var stateMachines map[string]*stateMachine

var (
	errStateMachineNotFound = errors.New("state machine not found")
	errInvalidState         = errors.New("invalid state for state machine")
	ErrInvalidMessage       = errors.New("invalid FTP message")
	ErrLoginFailed          = errors.New("login failed")
	ErrFrozenState          = errors.New("state is frozen")
)
