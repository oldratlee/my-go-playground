package internal

import "fmt"

// voidT is a value type and contains only one unique instance.
//
// Maybe use word `unit` or `singleton` is more precise,
// but `void` is more common and good enough.
type voidT = struct{}

// void is the only unique instance of voidT.
var void = voidT{} // = struct{}{}

func boxMessage(msg string) {
	bar := "=================================================="
	fmt.Printf("%v\n%v\n%[1]v\n", bar, msg)
}
