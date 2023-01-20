package internal

import "testing"

func Test_nilShowcase(t *testing.T) {
	/*
		defer func() {
			msg := recover()
			println("recovery:", msg)
		}()
	*/

	checkNilAndEmptySlice()

	checkNilAndEmptyMap()
}
