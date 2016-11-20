package main

import (
	"fmt"
	"time"
	"sync"
)

func sayGo(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 함수를 동기적으로 실행
	sayGo("Sync")

	// 함수를 비동기적으로 실행
	go sayGo("Async1")
	go sayGo("Async2")
	go sayGo("Async3")

	// 3초 대기
	fmt.Println("3초대기")
	time.Sleep(time.Second * 3)


	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() //Go루틴 모두 끝날 때까지 대기
}
