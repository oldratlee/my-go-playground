package internal

import (
	"bytes"
	"fmt"
	"unicode/utf8"
	"unsafe"
)

func commonCheckOfStringXInitLen() {
	boxMessage("stringInitAndCheckNil")

	var s0 string
	if s0 == "" {
		println("default/zero value of string EQUALS empty string")
	}
	fmt.Printf("len of default value of string is %v, rune count is %v\n", len(s0), utf8.RuneCountInString(s0))

	var s1 = "0123456789"
	var s2 = "hi"
	var s42 = "hi 🌞"

	fmt.Printf("\nlen: %2v, run count: %2v, string: '%v'\n", len(s1), utf8.RuneCountInString(s1), s1)
	fmt.Printf("len: %2v, run count: %2v, string: '%v'\n", len(s2), utf8.RuneCountInString(s2), s2)
	// 🌞 contains 1 rune, but contains 4 byte!
	fmt.Printf("len: %2v, run count: %2v, string: '%v'\n", len(s42), utf8.RuneCountInString(s42), s42)

}

func conversionsBetweenStringByteRuneShowcase() {
	boxMessage("conversionsBetweenStringByteRuneShowcase")

	////////////////////////////////////////////
	// string <=> []rune
	////////////////////////////////////////////
	s42 := "hi 🌞"
	rs := []rune(s42)
	fmt.Printf("string: '%v' => []rune: %+v\n", s42, rs)
	fmt.Printf("[]rune: %+v => string: '%v'\n", rs, string(rs))

	////////////////////////////////////////////
	// rune => string
	////////////////////////////////////////////
	var r rune = '🌞'
	fmt.Printf("\nrune: %c(%v) => string: '%v'\n", r, r, string(r))

	////////////////////////////////////////////
	// string <=> []byte
	////////////////////////////////////////////
	bs := []byte(s42)
	fmt.Printf("\nstring: '%v' => []byte: %+v\n", s42, bs)
	fmt.Printf("[]byte: %+v => string: '%v'\n", bs, string(bs))

	////////////////////////////////////////////
	// byte => string
	////////////////////////////////////////////
	var b byte = 'B'
	fmt.Printf("\nbyte: %c(%v) => string: '%v'\n", b, b, string(b))

	////////////////////////////////////////////
	// []rune <=> []byte
	////////////////////////////////////////////

	// indirect conversion: []rune => str => []bytes
	// TODO: how to directly converse??
	fmt.Printf("\n[]rune: %+v => []byte: %+v\n", rs, []byte(string(rs)))

	fmt.Printf("[]byte: %+v => []rune: %+v\n", bs, bytes.Runes(bs))
}

func stringIteration() {
	boxMessage("stringIteration")
	println("for range iteration return RUNE value with BYTE index, NOT return byte value")

	s42 := "hi 🌞!"

	println("\niterate rune value of string by for-range loop:")
	for i, r := range s42 {
		b := s42[i]
		fmt.Printf("index: %2v, rune: '%c'(%7v), byte value: %v\n",
			i, r, r, b)
	}

	println("\niterate with byte index only of string by for-range loop, lose bytes! NONSENSE!!")
	for i := range s42 {
		fmt.Printf("index: %2v, byte value: %v\n", i, s42[i])
	}

	println("\niterate with byte index only of string by for-index loop:")
	for i, l := 0, len(s42); i < l; i++ {
		fmt.Printf("index: %2v, byte value: %v\n", i, s42[i])
	}
}

func checkStringMemoryStruct() {
	boxMessage("checkStringMemoryStruct")

	var s0 string
	var s1 = "0123456789"
	var s2 = "hi"
	var s42 = "hi 🌞"

	var string2SizeOf = map[string]uintptr{
		s0:  unsafe.Sizeof(s0),
		s1:  unsafe.Sizeof(s1),
		s2:  unsafe.Sizeof(s2),
		s42: unsafe.Sizeof(s42),
	}

	for str, size := range string2SizeOf {
		fmt.Printf("sizeof: %v, string: '%v'\n", size, str)
	}
}
