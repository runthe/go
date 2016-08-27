package main

import (
	"math"
	"fmt"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}

	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}

/**
 if문 앞에 짧은 문장을 넣을 수 있다.  이부분 =>  v := math.Pow(x, n);
 if 블록 안에서만 사용가능하다.

  https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
 */