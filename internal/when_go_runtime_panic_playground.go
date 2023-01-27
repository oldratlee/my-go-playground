package internal

import (
	"fmt"
	"math/rand"
)

func arrayOperationPanic() {
	var a = [3]int{40, 41, 42}

	// out-of-bounds CONSTANT index, COMPILE ERROR!!
	//
	// println(a[4])
	// a[4] = 4

	// out-of-bounds index var(including negative index), runtime PANIC!!
	outOfBounds := rand.Intn(100) + 1000
	println(a[outOfBounds]) // PANIC!!
	a[outOfBounds] = 42     // PANIC!!
}

func sliceOperationPanic() {
	var a = []int{40, 41, 42}

	// out-of-bounds index var(including negative index), runtime PANIC!!
	println(a[42]) // PANIC!!
	a[4] = 42      // PANIC!!
}

func channelOperationPanic() {
	ch1 := make(chan int)
	close(ch1)
	ch2 := make(chan int)
	close(ch2)

	// Once a channel is closed, any attempts to write to the channel or close the channel again will panic.
	// Interestingly, attempting to read from a closed channel always succeeds.
	close(ch1) // PANIC!!
	ch2 <- 42  // PANIC!!

	// close a nil channel will panic.
	var nilChannel chan int // nil
	close(nilChannel)       // PANIC!!
}

func functionValueOperationCallNilFunctionValuePanic() {
	var f func() // nil

	f() // call a nil function value, PANIC!!
}

func callValueReceiverMethodWithNilPointReceiverPanic() {
	var nilPointFoo *fooStruct // nil

	nilPointFoo.pointerReceiverMethod() // OK

	nilPointFoo.valueReceiverMethod() // PANIC!!
}

type fooStruct struct {
}

//goland:noinspection GoMixedReceiverTypes
func (foo *fooStruct) pointerReceiverMethod() {
	println("call pointerReceiverMethod")
}

//goland:noinspection GoMixedReceiverTypes
func (foo fooStruct) valueReceiverMethod() {
	println("call valueReceiverMethod")
}

type fooInterface interface {
	pointerReceiverMethod()

	valueReceiverMethod()
}

func callMethodOfNilInterfacePanic() {
	var foo fooInterface // nil

	foo.pointerReceiverMethod() // PANIC!!

	foo.valueReceiverMethod() // PANIC!!
}

// callMethodOfInterfaceWithNilValuePanic:
// behaviour is same as callValueReceiverMethodUsingNilPointPanic.
func callMethodOfInterfaceWithNilValuePanic() {
	var nilPointFoo *fooStruct // nil point
	var foo fooInterface = nilPointFoo

	foo.pointerReceiverMethod() // OK

	foo.valueReceiverMethod() // PANIC!!
}

func pointerOperationDereferenceANilPointerPanic() {
	// dereference a nil pointer, PANIC!!
	var pi *int
	println(*pi) // PANIC!!
}

func typeAssertionsPanic() {
	type MyInt int
	var mi MyInt

	var i any
	i = mi

	i2 := i.(int) // PANIC!!
	// panic: interface conversion: interface {} is internal.MyInt, not int

	fmt.Printf("%v %T", i2, i2)
}
