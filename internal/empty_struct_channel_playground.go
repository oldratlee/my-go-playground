package internal

import (
	"fmt"
	"unsafe"
)

func createEmptyStructChannelWithElements() chan struct{} {
	c := make(chan voidT, 10)
	c <- void
	c <- void
	c <- void
	c <- void
	// it's ok to close before read all elements.
	close(c)

	return c
}

func emptyStructChannelReadChannelByCommaOk() {
	boxMessage("emptyStructChannelReadChannelByCommaOk")

	c := createEmptyStructChannelWithElements()

	for i := 0; ; i++ {
		v, ok := <-c
		if !ok {
			break
		}

		fmt.Printf("round %v, read: %v (size: %v | %v)\n",
			i, v, unsafe.Sizeof(v), unsafe.Sizeof(void))
	}
}
func emptyStructChannelReadByForRange() {
	boxMessage("emptyStructChannelReadByForRange")

	c := createEmptyStructChannelWithElements()

	var i int
	for v := range c {
		// no need to check whether the channel is closed by CommaOk pattern,
		// COOL!

		fmt.Printf("round %v, read: %v (size: %v | %v)\n",
			i, v, unsafe.Sizeof(v), unsafe.Sizeof(void))
		i++
	}
}
