package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 6}
	fmt.Println("p == ", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d]==%d\n", i, p[i])
	}
}

/**
   슬라이스는 배열의 값을 가리킴, 배열의 길이를 가짐
   []T는 T를 가지는 요소의 슬라이스이다

   https://www.facebook.com/groups/golangko/
   모든 내용은 http://go-tour-kr.appspot.com/
 */