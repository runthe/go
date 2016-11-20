package main

import "fmt"

type Rect struct {
	widht, height int
}

func (r Rect) area1() int {
	return r.widht * r.height
}

func (r *Rect) area2() int {
	r.widht++
	return r.widht * r.height
}

func main() {
	rect := Rect{10, 20}
	area := rect.area1() //메서드 호출
	fmt.Println(area)

	area2 := rect.area2()
	fmt.Println(area2)
}