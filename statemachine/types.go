package statemachine

import (
	"context"
	"errors"
)

type StateType int
type StateKey string
type StateHandler func(StateType, context.Context) context.Context
type State struct {
	tokens      [][]byte
	preHandler  StateHandler
	postHandler StateHandler
	nextState   StateType
}
type StateMachine struct {
	states map[StateType][]State
}

const Current StateKey = "CurrentState"

var (
	ErrInvalidStateContext = errors.New("invalid state context")
	ErrUndefinedState      = errors.New("transition not defined for this state")
	ErrNoTransition        = errors.New("no transition occurred")
)
