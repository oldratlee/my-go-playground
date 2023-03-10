package internal

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// code from "CHAPTER 10 Concurrency in Go"
// of "Learning Go: An Idiomatic Approach to Real-World Go Programming"
func operateMultiplyChannelWithSelectStatementAvoidStarvation() {
	boxMessage("operateMultiplyChannelWithSelectStatementAvoidStarvation")

	ch1 := make(chan int)
	ch2 := make(chan int)

	// start a sub-goroutine
	go func() {
		v1 := <-ch1
		fmt.Printf("read %v from channel 1 in sub-goroutine\n", v1)
		// NOTE: write channel 1 then read channel 2 SEQUENTIALLY!!
		ch2 <- 42
		println("[operateMultiplyChannelWithSelectStatementAvoidStarvation] sub-goroutine exit(NO goroutine leak)")
	}()

	// NOTE: read channel 1 and write channel 2 by select statement, NOT sequentially.
	//
	// if write/read channel use sequential code logic as above
	// instead of select statement, cause STARVATION!!
	//
	// for this simple test code(only 2 goroutines), cause DEADLOCK:
	//   fatal error: all goroutines are asleep - deadlock!
	//
	// example SEQUENTIAL codes:
	//
	// v2 := <-ch2
	// fmt.Printf("read %v from channel 2 in main goroutine\n", v2)
	// ch1 <- 21

	for i := 0; i < 2; i++ {
		// run select statement twice
		select {
		case v2 := <-ch2:
			fmt.Printf("read %v from channel 2 in main goroutine\n", v2)
		case ch1 <- 21:
		}
	}

	// sleep to check whether sub-goroutine exited
	time.Sleep(5 * time.Millisecond)
}

func _search(s string) []string {
	_randomSleep()

	const count = 3
	data := make([]string, count)
	for i := 0; i < count; i++ {
		data[i] = fmt.Sprintf("%s/%d", s, i)
	}
	return data
}

type _searcherT = func(string) []string

// code from "CHAPTER 10 Concurrency in Go"
// of "Learning Go: An Idiomatic Approach to Real-World Go Programming"
func doneChannelPatternAkaAChannelOnlyUsedToPublishCloseEvent() {
	boxMessage("doneChannelPatternAkaAChannelOnlyUsedToPublishCloseEvent")

	done := make(chan voidT)
	result := make(chan []string)

	var searchers = [...]_searcherT{_search, _search, _search}
	for i, s := range searchers {
		go func(searcher _searcherT, num int) {
			id := fmt.Sprintf("foo%d", num)

			select {
			case result <- searcher(id):
				fmt.Printf("[doneChannelPatternAkaAChannelOnlyUsedToPublishCloseEvent] searcher goroutine %s put result and exit(NO goroutine leak)\n", id)
			case <-done:
				fmt.Printf("[doneChannelPatternAkaAChannelOnlyUsedToPublishCloseEvent] searcher goroutine %s is canceled by done channnel and exit(NO goroutine leak)\n", id)
			}

		}(s, i)
	}

	// Got one result from a searcher goroutine
	r := <-result

	// close done channel, NOTIFY other searcher goroutine to exit
	close(done)

	fmt.Printf("Got one result: %+v\n", r)

	// sleep to check whether all searcher goroutine exited
	time.Sleep(10 * time.Millisecond)
}

func _countTo(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan voidT)
	cancel := func() {
		close(done)
	}

	// start a goroutine to fill the int sequence channel
	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				println("goroutine of _countTo is CANCELED")
				return
			case ch <- i:
				println("goroutine of _countTo produce", i)
			}
		}
		close(ch)
		println("goroutine of _countTo is CLOSED")
	}()

	return ch, cancel
}

// code from "CHAPTER 10 Concurrency in Go"
// of "Learning Go: An Idiomatic Approach to Real-World Go Programming"
func cancelFunctionPatternToTerminateAGoroutine() {
	boxMessage("cancelFunctionPatternToTerminateAGoroutine")

	ch, cancel := _countTo(4)
	defer cancel()

	for i := range ch {
		println("main read ", i)

		r := 2 + rand.Intn(2) // random value of 2 or 3
		if i >= r {
			break
		}
	}

	// sleep to check whether sub-goroutine exited
	time.Sleep(10 * time.Millisecond)
}

// code from "CHAPTER 10 Concurrency in Go"
// of "Learning Go: An Idiomatic Approach to Real-World Go Programming"
func nilifyTheCaseChannelPatternToTurningOffACaseInASelect() {
	boxMessage("nilifyTheCaseChannelPatternToTurningOffACaseInASelect")

	// prepare the input channel
	const dataCount = 3
	in1 := make(chan int, dataCount)
	in2 := make(chan int, dataCount)
	// fill mock data
	for i := 0; i < dataCount; i++ {
		in1 <- i + 100
		in2 <- i + 200
	}
	close(in1)
	close(in2)

	// prepare the done channel
	done := make(chan voidT)
	time.AfterFunc(5*time.Millisecond, func() {
		close(done)
	})

	// main for-select loop
	for {
		select {
		case v, ok := <-in1:
			if !ok {
				in1 = nil // the case will never succeed again!
				println("NILIFY the closed channel in1")
				continue
			}
			// process the v that was read from in1
			println("read from in1:", v)
		case v, ok := <-in2:
			if !ok {
				in2 = nil // the case will never succeed again!
				println("NILIFY the closed channel in2")
				continue
			}
			// process the v that was read from in2
			println("read from in2:", v)
		case <-done:
			println("CANCELLED, stop main for-select loop")
			return
		}
	}
}

func _doSomeWork() (int, error) {
	_randomSleep()

	if rand.Intn(2) > 0 {
		println("[_doSomeWork] do some work error")
		return 0, errors.New("do some work error")
	} else {
		println("[_doSomeWork] do some work success")
		return 42, nil
	}
}

func _timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan voidT)

	// start a sub-goroutine to do some work
	go func() {
		result, err = _doSomeWork()
		close(done)
	}()

	// time limit for work
	select {
	case <-done:
		fmt.Printf("[_timeLimit] work finished: %d, %#v\n", result, err)
		return result, err
	case <-time.After(7500 * time.Microsecond):
		println("[_timeLimit] work timed out")
		return 0, errors.New("work timed out")
	}
}

// code from "CHAPTER 10 Concurrency in Go"
// of "Learning Go: An Idiomatic Approach to Real-World Go Programming"
func timeAfterCasePatternToTimeOutCode() {
	boxMessage("timeAfterCasePatternToTimeOutCode")

	result, err := _timeLimit()
	fmt.Printf("[main] work return with time limit: %d, %#v\n", result, err)

	// sleep to check whether all sub-goroutine exited
	time.Sleep(10 * time.Millisecond)
}

// _randomSleep: sleep about 10ms(5ms ~ 10ms)
func _randomSleep() {
	d := 5 + rand.Int63n(5)
	time.Sleep(time.Duration(d) * time.Millisecond)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
