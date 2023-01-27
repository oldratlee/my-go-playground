package internal

import "testing"

func Test_arrayOperationPanic(t *testing.T) {
	arrayOperationPanic()
}

func Test_sliceOperationPanic(t *testing.T) {
	sliceOperationPanic()
}

func Test_channelOperationPanic(t *testing.T) {
	t.Run("", func(t *testing.T) {
		channelOperationPanic()
	})
}

func Test_functionValueOperationCallNilFunctionValuePanic(t *testing.T) {
	functionValueOperationCallNilFunctionValuePanic()
}

func Test_callValueReceiverMethodWithNilPointReceiverPanic(t *testing.T) {
	callValueReceiverMethodWithNilPointReceiverPanic()
}

func Test_callMethodOfNilInterfacePanic(t *testing.T) {
	callMethodOfNilInterfacePanic()
}

func Test_callMethodOfInterfaceWithNilValuePanic(t *testing.T) {
	callMethodOfInterfaceWithNilValuePanic()
}

func Test_pointerOperationDereferenceANilPointerPanic(t *testing.T) {
	pointerOperationDereferenceANilPointerPanic()
}

func Test_typeAssertionsPanic(t *testing.T) {
	typeAssertionsPanic()
}
