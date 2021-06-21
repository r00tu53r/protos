package fsm

import (
	"strings"
)

func init() {
	stateMachines = make(map[string]*stateMachine)
	lsm := newLoginStateMachine()
	stateMachines[User] = lsm
	stateMachines[Password] = lsm
	stateMachines[Account] = lsm
}

func GetStateMachine(cmd string) (*stateMachine, error) {
	_, ok := Commands[strings.ToUpper(cmd)]
	if !ok {
		return nil, errStateMachineNotFound
	}
	return stateMachines[strings.ToUpper(cmd)], nil
}

func newStateMachine() *stateMachine {
	return &stateMachine{
		states: make(map[State][]fsmState),
	}
}

func (sm *stateMachine) addState(currentState State,
	tokens []string, nextState State, isRequest bool) {
	states, _ := sm.states[currentState]
	newState := fsmState{
		tokens:    tokens,
		next:      nextState,
		isRequest: isRequest,
	}
	states = append(states, newState)
	sm.states[currentState] = states
}

func (sm *stateMachine) Transition(state *FSMState, line string) (*FSMState, error) {
	noop := true
	if state.Current == StateSuccess {
		return state, nil
	}
	if state.Current == StateError || state.Current == StateFailed {
		return state, ErrLoginFailed
	}
	fsmStates, ok := sm.states[state.Current]
	if !ok {
		return nil, errInvalidState
	}
	for _, fs := range fsmStates {
		if fs.isRequest == state.IsRequest {
			for _, tok := range fs.tokens {
				if state.IsRequest && state.ReqOp == tok {
					state.Current = fs.next
					noop = false
					break
				}
				if !state.IsRequest && state.RespCode == tok {
					state.Current = fs.next
					noop = false
					break
				}
			}
		}
		if !noop {
			break
		}
	}
	if noop {
		return state, ErrFrozenState
	}
	return state, nil
}

func (s *FSMState) Final() bool {
	if s.Current == StateSuccess ||
		s.Current == StateFailed || s.Current == StateError {
		return true
	}
	return false
}
