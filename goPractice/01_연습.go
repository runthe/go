package main

import "fmt"

func init() {
	fmt.Println("init")
}

func main() {

	memory()
	fmt.Println("======")
	instance()
	defer exit()
}

func inc(i *int) {
	*i = *i + 1
}

func exit() {
	fmt.Println("exit")
}

func memory() {
	number := 5
	inc(&number)
	fmt.Println(number)

	var p *int
	var pp **int

	ii := 1
	p = &ii
	pp = &p
	fmt.Println(ii, *p, **pp)

	var po = new(int)
	*po = 11

	fmt.Println(po)
	fmt.Println(*po)
}

func instance() {
	//m1 := make(map[string]string, 5)
	m2 := make(map[string]string)
	m2["one"] = "1"

	m3 := map[string]string{}
	m4 := map[string]string{"one":"1", "two":"2"}

	fmt.Println(m2, m3, m4)

}