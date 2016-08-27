package main

import "fmt"

func main() {

	// 'var' 는 한개나 여러 개의 변수를 선언한다
	var a string = "initial"
	fmt.Println(a)

	//동시에 변수 선언
	var b, c int = 3, 5
	fmt.Println(b, c)

	//Go는 변수가 초기화될떄 타입을 추론한다.
	var d = true
	fmt.Println(d)

	//초기화값이 없는 변수는 zero-value. 숫자의 경우는0
	var e int
	fmt.Println(e)

	var f string
	fmt.Println(f)

	// := 는 초기화와 선언의 요약이다 로컬에 쓰임
	// g := short var g string = short
	g := "short"
	fmt.Println(g)

}
