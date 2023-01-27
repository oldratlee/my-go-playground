package internal

import "testing"

func Test_channelSelectOperationAvoidStarvation(t *testing.T) {
	operateMultiplyChannelWithSelectStatementAvoidStarvation()
}

func Test_doneChannelPattern(t *testing.T) {
	doneChannelPatternAkaAChannelOnlyUsedToPublishCloseEvent()
}

func Test_cancelFunctionPatternToTerminateAGoroutine(t *testing.T) {
	cancelFunctionPatternToTerminateAGoroutine()
}
