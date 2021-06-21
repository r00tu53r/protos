package fsm

func newLoginStateMachine() *stateMachine {
	sm := newStateMachine()
	sm.addState(StateBegin, []string{User}, StateWaitForReply, true)
	sm.addState(StateWaitForReply, []string{"331"}, StateNeedPassword, false)
	sm.addState(StateNeedPassword, []string{Password}, StateWaitForReply, true)
	sm.addState(StateWaitForReply, []string{"230"}, StateSuccess, false)
	sm.addState(StateWaitForReply, []string{"530"}, StateFailed, false)
	return sm
}
