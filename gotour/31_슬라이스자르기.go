package main

import "fmt"

func main() {
	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p == ", p)
	fmt.Println("p[2:3] == ", p[2:3])

	//missing low index implies 0
	fmt.Println("p[:3] == ", p[:3])

	//missing high index implies len(s)
	fmt.Println("p[4:]", p[4:])
}

/**
슬라이스에서 슬라이스 이름[시작:끝1]로 재분할가능하다
예 저기위에서 슬라이스 이름[2:3]은 2번째 5가 나온다

 https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
 */
