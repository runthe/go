package main

import "fmt"

type Vertext struct {
	X, Y int
}

func main() {
	fmt.Println(Vertext{1, 2})
}

/**
 구조체는 필드와 데이터들의 조합이다
 type 선언으로 struct의 이름을 정할수 있다

 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
 */