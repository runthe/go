package main

import "fmt"

type Vertex struct {
	X, Y int
}

func main() {
	v := new(Vertex)
	fmt.Println(v)

	v.X, v.Y = 11, 9
	fmt.Println(v)

	var vv *Vertex = new(Vertex)
	vv.X, vv.Y = 0,9
	fmt.Println(vv)

	var vl = Vertex{}
	vl.X, vl.Y = 1,2
	fmt.Println(vl)

//	var vvv Vertex = new(Vertex)
//	vvv.X, vvv.Y = 1,2
//	fmt.Println(vvv)
}

/**
new 는 모든 필드가 0이 할당된 T타입의 포인터를 반환함
var t *T = new(T) 또는 t := new(T)로 할당한다

https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
 */