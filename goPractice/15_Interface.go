package main

import (
	"math"
	"fmt"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Rect struct {
	width, height float64
}

type Circle struct {
	radius float64
}


//Rect 타입에 대한 Shape인터페이스 구현
func (r Rect) area() float64 {
	return r.width * r.height
}

func (r Rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

//Circle 타입에 대한 Shape인터페이스 구현
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {
	r := Rect{10., 20.}
	c := Circle{10}

	showArea(r, c)

	var x interface{}
	x = 1
	x = "Tom"

	printIt(x)

	var a interface{} = 1

	//a와 i는 다이나믹한 type 값은 1
	i := a
	//j는 int타입, 값은 1
	j := a.(int)

	//포인터주소 출력
	println(i)
	//1출력
	println(j)

}

func printIt(v interface{}) {
	fmt.Println(v)
}

func showArea(shapes ...Shape) {
	for _, s := range shapes {
		//인터페이스 메서드 호출
		a := s.area()
		println(a)
	}
}

