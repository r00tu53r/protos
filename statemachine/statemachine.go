package statemachine

import (
	"bytes"
	"context"
)

func New() *StateMachine {
	return &StateMachine{
		states: make(map[StateType][]State),
	}
}

func GetState(ctx context.Context) (*StateType, error) {
	current, ok := ctx.Value(Current).(StateType)
	if ok {
		return &current, nil
	}
	return nil, ErrInvalidStateContext
}

func SetState(
	ctx context.Context,
	startState StateType) context.Context {
	var uCtx context.Context
	if ctx == nil {
		uCtx = context.Background()
	}
	uCtx = context.WithValue(ctx, Current, startState)
	return uCtx
}

func (sm *StateMachine) AddState(
	current StateType,
	preHandler StateHandler,
	next StateType,
	postHandler StateHandler,
	tokens [][]byte) {

	states, _ := sm.states[current]
	newState := State{
		tokens:      tokens,
		nextState:   next,
		preHandler:  preHandler,
		postHandler: postHandler,
	}
	states = append(states, newState)
	sm.states[current] = states
}

func (sm *StateMachine) Transition(
	stateContext context.Context,
	payload []byte) (context.Context, error) {

	var noop bool

	if stateContext == nil {
		return stateContext, ErrInvalidStateContext
	}
	current, ok := stateContext.Value(Current).(StateType)
	if !ok {
		return stateContext, ErrInvalidStateContext
	}
	fsmStates, ok := sm.states[current]
	if !ok {
		return stateContext, ErrUndefinedState
	}

	noop = true
	for _, fs := range fsmStates {
		for _, tok := range fs.tokens {
			if bytes.Compare(payload, tok) == 0 {
				if fs.preHandler != nil {
					stateContext = fs.preHandler(current, stateContext)
				}
				stateContext = context.WithValue(stateContext, Current, fs.nextState)
				if fs.postHandler != nil {
					stateContext = fs.postHandler(fs.nextState, stateContext)
				}
				noop = false
				break
			}
		}
		if !noop {
			break
		}
	}
	if noop {
		return stateContext, ErrNoTransition
	}
	return stateContext, nil
}
