package main

import "fmt"

func main() {
	//슬라이스 변수 선언
	var a []int
	a = []int{1, 2, 3 } //슬라이스에 리터럴값 저장
	a[1] = 10
	fmt.Println(a)

	//첫번째 타입 두번쨰 length, 세번쨰 내부 배열의 최대 길이 지정 모든요소가 Zero Value인 슬라이스
	//세번째 value를 생략하면 length와 같은 값을 갖는다
	s := make([]int, 5, 10)
	fmt.Println(s)
	fmt.Println(len(s), cap(s))

	var t []int
	println(len(t), cap(t))

	tt := []int{0, 1, 2, 3, 4, 5}
	tt = tt[2:5]
	fmt.Println(tt)

	ss := []int{0, 1}
	ss = append(ss, 2)
	ss = append(ss, 3, 4, 5)
	fmt.Println(ss)

	sliceA := make([]int, 0, 3)
	for i := 1; i < 15; i++ {
		sliceA = append(sliceA, i)

		fmt.Println(len(sliceA), cap(sliceA))
	}

	sliceB := []int{4, 5, 6}
	sliceA = append(sliceA, sliceB...)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA) // [1 2 3 4 5 6] 출력

	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)
	copy(target, source)
	fmt.Println(target)  // [0 1 2 ] 출력
	println(len(target), cap(target)) // 3, 6 출력

}