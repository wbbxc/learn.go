package main

import (
	"fmt"
)

func test(a int){
	a++
	fmt.Println(a)
}


func main() {
	a:=2

	for i:=0;i<3;i++{

		test(a)

	}
	test(a)


}
