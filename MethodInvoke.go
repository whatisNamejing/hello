package main

import "fmt"

func main() {
	x,y:=toAll(1,"xx")
	fmt.Println(x,y)

	//返回值顺序接受
	var a,b = toAll(200,"可以啊")
	fmt.Println(a,b)
}
/**
	多返回
 */
func toAll(x int,y string) (int,string) {
	return x,y
}