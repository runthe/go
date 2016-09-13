/*
	author twitter : @mmcgrana
	https://gobyexample.com/arrays

	Go by Example: Arrays

   Go에서 배열은 특정 길이를 가진, 번호가 붙은 요소들의 나열이다.
*/
package main

import "fmt"

func main() {

	var a [5]int
	fmt.Println(a)

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println(len(a))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	fmt.Println(twoD)
}
