package main

import (
	"os"
	"fmt"
)

func main() {
	openFile("1.txt")
	println("Done") // 이 문장 실행됨
}

func openFile(fn string) {
	// defere 함수. panic 호출시 실행됨
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	// 파일 close 실행됨
	defer f.Close()
}

func test() {
	defer fmt.Println("1")
	defer fmt.Println("2")

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("ERROR", r)
		}
	}()

	defer fmt.Println("3")

	panic("으아")

	defer fmt.Println("4")
}

