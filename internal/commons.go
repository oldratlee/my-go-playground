package internal

// voidT is a value type and only contains one instance.
//
// Maybe use word `unit` or `singleton` is more precise,
// but `void` is more common and good enough.
type voidT = struct{}

// void is the only unique instance of voidT.
var void = voidT{} // = struct{}{}
