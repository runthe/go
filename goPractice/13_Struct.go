package main

import "fmt"

type person struct {
	name string
	age  int
}

type dict struct {
	data map[int]string
}

func newDict() *dict {
	d := dict{}
	d.data = map[int]string{}
	return &d //포인터 전달
}

func main() {
	//person 객체 생성
	p := person{}

	//필드값 설정
	p.name = "Lee"
	p.age = 10

	fmt.Println(p)

	p1 := person{"bob", 20}
	p2 := person{name:"haha", age:50}

	fmt.Println(p1)
	fmt.Println(p2)

	p3 := new(person)
	p3.name = "Lee"

	fmt.Println(p3)

	dic := newDict() //생성자 호출
	dic.data[1] = "A"
}