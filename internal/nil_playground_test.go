package internal

import "testing"

func Test_nilOrEmptyShowcase(t *testing.T) {
	/*
		defer func() {
			msg := recover()
			println("recovery:", msg)
		}()
	*/

	checkNilOrEmptySlice()

	checkNilOrEmptyMap()
}
