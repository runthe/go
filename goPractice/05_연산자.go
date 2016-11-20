package main

import (
	"fmt"
)

func main() {

	var k int = 10
	var p = &k  //k의 주소를 할당

	fmt.Println(p)
	fmt.Println(*p)
}
