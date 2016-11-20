package main

import "fmt"


//type문 원형정의
type calculator func(x int, y int) int

func main() {

	sum := func(numbers ...int) {
		for index, number := range numbers {
			fmt.Println(index, number)
		}
	}

	sum(1, 5, 7, 9, 15)

	add := func(i int, j int) int {
		return i + j
	}

	//익명함수를 변수에 할당
	r1 := calc(add, 10, 20)
	fmt.Println(r1)

	//익명함수를 직접 할당
	r2 := calc(func(x int, y int) int {
		return x - y
	}, 10, 20)
	fmt.Println(r2)

	next := nextValue()

	fmt.Println(next())
	fmt.Println(next())
	fmt.Println(next())
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func calc2(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

func nextValue() func() int {
	i := 0;

	return func() int {
		i++
		return i
	}
}