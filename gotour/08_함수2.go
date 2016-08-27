package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(52, 10))
}

/*
 두개 이상의 매개변수가 같은 타입일때 타입을 취하는 마지막 매개변수 타입만
 명확하게 명시 예르들어 x int, y int x y int 로 생략

https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
 */