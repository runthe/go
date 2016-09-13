/*
    author twitter : @mmcgrana
	https://gobyexample.com/tickers
    Go by Example: tickers

	timer는 당신이 미래에 어떤 일을 하고 싶을 때 쓰고 ticker는 당신이 정기적으로
	어떤 일을 반복적으로 하기 원할 때 쓴다. 여기 우리가 멈출 때까지 정기적으로 tick하는
	ticker 예제가 있다.  (틱장애가 떠오른다...역주)

*/

package main

import (
	"time"
	"fmt"
)

func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	go func() {
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1500)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
