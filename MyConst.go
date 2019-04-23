package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var a = "ssa"
	var s = len(a)
	var b = unsafe.Sizeof(a)
	fmt.Println(s,b)
}
