package internal

import "fmt"

func checkNilAndEmptySlice() {
	boxMessage("checkNilSlice")

	////////////////////////////////////////////
	// nil slice
	////////////////////////////////////////////

	var nilSlice []int
	// append elements to nil slice, OK!
	s1 := append(nilSlice, 42, 43)
	fmt.Printf("append to nil slice is OK, slice is %+v\n", s1)
	if nilSlice == nil {
		println("nilSlice is nil")
	}

	////////////////////////////////////////////
	// empty slice
	////////////////////////////////////////////

	var emptySlice = []int{}
	// append elements to empty slice, OK!
	//
	// the behavior of `append`ing to empty slice is same as `nil`,
	// so is not recommend to use `empty slice` in most case.
	//
	// and it's a WARNING in `GoLang`.
	s2 := append(emptySlice, 420, 430)
	fmt.Printf("append to empty slice is OK, slice is %+v\n", s2)

	if emptySlice == nil {
		println("empty slice EQUALS nil")
	} else {
		println("empty slice NOT equals nil")
	}
}

func checkNilAndEmptyMap() {
	boxMessage("checkNilAndEmptyMap")

	////////////////////////////////////////////
	// empty map
	////////////////////////////////////////////

	var emptyMap = map[int]int{}
	// write elements to empty map, OK!
	emptyMap[1] = 42
	fmt.Printf("write to empty map is OK, map is %+v\n", emptyMap)

	////////////////////////////////////////////
	// nil map
	////////////////////////////////////////////

	var nilMap map[int]int
	// write elements to empty map, PANIC!
	nilMap[1] = 420
	// below line code never run!
	fmt.Printf("write to nil map, PANIC, map is %+v\n", nilMap)
}
