package internal

import "testing"

func Test_checkNilOrEmptySlice(t *testing.T) {
	checkNilOrEmptySlice()
}

func Test_checkNilOrEmptyMap(t *testing.T) {
	/*
		defer func() {
			msg := recover()
			println("recovery:", msg)
		}()
	*/

	checkNilOrEmptyMap()
}

func Test_checkNilLambda(t *testing.T) {
	checkNilLambda()
}
