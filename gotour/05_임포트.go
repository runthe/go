package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Nextafter(2, 3))
}


/**
 내용: 여러개의 package를 소괄호로 감싸서 임포트로 표현함을 해보자
 import("fmt" "math")
 https://www.facebook.com/groups/golangko/
 모든 내용은 http://go-tour-kr.appspot.com/
 */
