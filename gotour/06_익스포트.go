//패키지명은 대문자를 써야 외부참조 가능!
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.pi)
	//fmt.Println(math.Pi)
}

/**
 내용 : pi를 Pi로 변경해야 한다 외부참조시에 대문자로 쓰여져 exported된것들만 접근가능하다
  https://www.facebook.com/groups/golangko/
모든 내용은 http://go-tour-kr.appspot.com/
 */
