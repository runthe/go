package main

import (
	"fmt"
	"math"
)

// const 는 상수를 선언한다

const s string = "constant"

func main() {
	fmt.Println(s)

	//const 는 var가 나타날 수 잇는 어디든지 등장할 수 있따

	//상수표현식은 임의적인 정확도의 연산을 수행한다
	const n = 5000000000
	const d = 3e20 / n
	fmt.Println(d)

	//숫자상수는 명시적형변환 같은 타입이 주어질떄까지 어떤 타입도 갖지 않는다.
	fmt.Println(int64(d))

	//숫잦는 그것을 필요로 하는 맥락속에서 변수할당이나 함수호출로써 사용되므로써 타입이 주어
	//질떄 바뀐다 예를 들면 여기서 math.sin은 float64를 필요로 한다
	fmt.Println(math.Sin(n))

}
