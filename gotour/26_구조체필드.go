package main

import "fmt"

type Vertext struct {
	X, Y int
}

func main() {
	v := Vertext{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

/**
 구조체 필드는 . 으로 접근한다
 */
