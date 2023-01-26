package internal

import "fmt"

func createChannelWithIntElements() <-chan int {
	ch := make(chan int, 10)
	// write 5 int elements
	for i := 0; i < 5; i++ {
		ch <- i
	}
	// it's ok to close before read (all) elements.
	//
	// in this showcase,
	// if created channel is not closed before read, cause DEADLOCK:
	//   fatal error: all goroutines are asleep - deadlock!
	close(ch)

	return ch
}

func readChannelByCommaOkPattern() {
	boxMessage("readChannelByCommaOkPattern")

	ch := createChannelWithIntElements()

	for i := 0; ; i++ {
		v, ok := <-ch
		if !ok {
			break
		}

		fmt.Printf("round %v, read: %v\n", i, v)
	}
}

func readChannelByForRange() {
	boxMessage("readChannelByForRange")

	ch := createChannelWithIntElements()

	i := 0
	for v := range ch {
		// no need to check whether the channel is closed by CommaOk pattern,
		// concise and COOL!

		fmt.Printf("round %v, read: %v\n", i, v)
		i++
	}
}
