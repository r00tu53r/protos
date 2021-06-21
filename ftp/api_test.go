package ftp

import (
	"testing"

	"github.com/r00tu53r/protos/ftp/fsm"
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
	state, err := Run(ftpLoginMessages[0], nil)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateWaitForReply, "must be equal")

	state, err = Run(ftpLoginMessages[1], state)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateNeedPassword, "must be equal")

	state, err = Run(ftpLoginMessages[2], state)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateWaitForReply, "must be equal")

	state, err = Run(ftpLoginMessages[3], state)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateSuccess, "must be equal")
}

func TestFailedLogin(t *testing.T) {
	// A series of message exchanges between the FTP
	// client and the server for a successful login.
	ftpLoginMessages := []string{
		"USER test",
		"331 User test OK. Password required.",
		"PASS wrongpassword",
		"530 Login authentication failed",
	}
	state, err := Run(ftpLoginMessages[0], nil)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateWaitForReply, "must be equal")

	state, err = Run(ftpLoginMessages[1], state)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateNeedPassword, "must be equal")

	state, err = Run(ftpLoginMessages[2], state)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateWaitForReply, "must be equal")

	state, err = Run(ftpLoginMessages[3], state)
	assert.Equal(t, err, nil, "error must be nil")
	assert.Equal(t, state.Current, fsm.StateFailed, "must be equal")
}
