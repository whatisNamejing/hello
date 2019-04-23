package main

import "fmt"

func main() {
	for i:=1;i<=20;i++{
		fmt.Println()
		for j:=1;j<=i;j++{
			fmt.Print("*")
		}
	}
	var fileName string
	fmt.Scanln(&fileName)
	if fileName=="1"{
		fmt.Println("成功")
	}else if fileName=="2" {
		fmt.Println("垃圾")
	}else{
		fmt.Println("不行啊")
	}
}
