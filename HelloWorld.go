package main

import "fmt"

func main() {
	const (
		a = iota
		b = "ss"
		c = iota
	)
	fmt.Println("hello world !",a,b,c)
}
