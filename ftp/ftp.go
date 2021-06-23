package ftp

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/r00tu53r/protos/statemachine"
)

const (
	StateBegin statemachine.StateType = iota
	StateWaitForReply
	StateNeedPassword
	StateSuccess
	StateFailed
)

const CommandKey statemachine.StateKey = "FTPCommand"

var (
	ErrInvalidMessage       = errors.New("invalid FTP message")
	ErrStateMachineNotFound = errors.New("state machine not found")
)

// TODO track the session with IP port tuples, PORT commands
func Run(message string, ctx context.Context) (context.Context, error) {
	if message == "" {
		return nil, ErrInvalidMessage
	}
	reqOp, isReq := IsRequest(message)
	respCode, isResp := isResponse(message)
	if !isReq && !isResp {
		return nil, ErrInvalidMessage
	}
	_, err := statemachine.GetState(ctx)
	if err != nil {
		return nil, err
	}
	if isReq {
		ctx = context.WithValue(ctx, CommandKey, reqOp)
	}
	reqOp, ok := ctx.Value(CommandKey).(string)
	if !ok {
		return nil, ErrStateMachineNotFound
	}
	sm, err := getStateMachine(reqOp)
	if err != nil {
		return nil, err
	}
	if isReq {
		ctx, err = sm.Transition(ctx, []byte(reqOp))
	} else {
		ctx, err = sm.Transition(ctx, []byte(strconv.Itoa(respCode)))
	}
	return ctx, err
}

func IsRequest(line string) (string, bool) {
	firstWord := strings.Split(line, " ")[0]
	_, ok := Commands[strings.ToUpper(firstWord)]
	return firstWord, ok
}

func getCode(word string) (int, error) {
	last := len(word) - 1
	code, err := strconv.Atoi(word)
	if err != nil {
		if word[last] == '-' {
			code, err = strconv.Atoi(word[:last])
			if err != nil {
				return -1, err
			}
			return code, nil
		}
		return -1, err
	}
	return code, nil
}

func isResponse(line string) (int, bool) {
	firstWord := strings.Split(line, " ")[0]
	code, err := getCode(firstWord)
	if err != nil {
		return -1, false
	}
	return code, true
}
