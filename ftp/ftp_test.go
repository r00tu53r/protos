package ftp

import (
	"context"
	"testing"

	"github.com/r00tu53r/protos/statemachine"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	// A series of message exchanges between the FTP
	// client and the server for a successful login.
	ftpLoginMessages := []string{
		"USER test",
		"331 User test OK. Password required.",
		"PASS hello",
		"230 OK. Current directory is /",
	}

	ctx := statemachine.SetState(context.Background(), StateBegin)

	ctx, err := Run(ftpLoginMessages[0], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err := statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateWaitForReply, "must be equal")

	ctx, err = Run(ftpLoginMessages[1], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err = statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateNeedPassword, "must be equal")

	ctx, err = Run(ftpLoginMessages[2], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err = statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateWaitForReply, "must be equal")

	ctx, err = Run(ftpLoginMessages[3], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err = statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateSuccess, "must be equal")
}

func TestLoginFailed(t *testing.T) {
	// A series of message exchanges between the FTP
	// client and the server for a failed login.
	ftpLoginMessages := []string{
		"USER test",
		"331 User test OK. Password required.",
		"PASS hello",
		"530 Login authentication failed",
	}

	ctx := statemachine.SetState(context.Background(), StateBegin)

	ctx, err := Run(ftpLoginMessages[0], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err := statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateWaitForReply, "must be equal")

	ctx, err = Run(ftpLoginMessages[1], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err = statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateNeedPassword, "must be equal")

	ctx, err = Run(ftpLoginMessages[2], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err = statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateWaitForReply, "must be equal")

	ctx, err = Run(ftpLoginMessages[3], ctx)
	assert.Equal(t, err, nil, "error must be nil")
	currState, err = statemachine.GetState(ctx)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, *currState, StateFailed, "must be equal")
}
