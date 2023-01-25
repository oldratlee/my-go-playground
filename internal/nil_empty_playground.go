package internal

import "fmt"

func checkNilOrEmptySlice() {
	boxMessage("checkNilOrEmptySlice")

	////////////////////////////////////////////
	// nil slice
	////////////////////////////////////////////

	var nilSlice []int
	// nilSlice is nil :
	if nilSlice == nil {
		println("nilSlice is nil")
	}
	fmt.Printf("len of nilSlice(%+v) is %v\n", nilSlice, len(nilSlice))

	// append elements to nil slice, OK!
	// return a new slice with appended elements
	s1 := append(nilSlice, 42, 43)
	fmt.Printf("append to nil slice is OK, return result is %+v\n", s1)

	////////////////////////////////////////////
	// empty slice
	////////////////////////////////////////////

	var emptySlice = []int{}
	fmt.Printf("len of emptySlice(%+v) is %v\n", emptySlice, len(emptySlice))

	// append elements to empty slice, OK!
	//
	// the behavior of `append`ing to empty slice is same as `nil`,
	// so is not recommend to use `empty slice` in most case.
	//
	//
	// and it's a *WARNING* in `GoLand` IDE:
	//
	// Reports slice declarations with empty literal initializers used instead of nil.
	// An empty slice can be represented by nil or an empty slice literal.
	// They are functionally equivalent — their len and cap are both zero — but the nil slice is the preferred style.
	// For more information about empty slices,
	// see [Declaring Empty Slices at golang.org](https://golang.org/wiki/CodeReviewComments#declaring-empty-slices).
	//
	s2 := append(emptySlice, 420, 430)
	fmt.Printf("append to empty slice is OK, return result is %+v\n", s2)

	// empty slice NOT equals nil.
	if emptySlice == nil {
		println("empty slice EQUALS nil")
	} else {
		println("empty slice NOT equals nil")
	}
}

func checkNilOrEmptyMap() {
	boxMessage("checkNilOrEmptyMap")

	////////////////////////////////////////////
	// empty map
	////////////////////////////////////////////

	var emptyMap = map[int]int{}
	fmt.Printf("len of emptyMap(%+v) is %v\n", emptyMap, len(emptyMap))

	// write elements to empty map, OK and success!
	emptyMap[1] = 42
	fmt.Printf("write to empty map is OK, map becomes %+v\n", emptyMap)

	////////////////////////////////////////////
	// nil map
	////////////////////////////////////////////

	var nilMap map[int]int
	// nilMap is nil :
	if nilMap == nil {
		println("nilMap is nil")
	}
	fmt.Printf("len of nilMap(%+v) is %v\n", nilMap, len(nilMap))

	// write elements to empty map, PANIC!
	nilMap[1] = 420
	// below line code never run!
	fmt.Printf("write to nil map, PANIC, map becomes %+v\n", nilMap)
}