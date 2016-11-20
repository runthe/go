package main

import "fmt"

func main() {
	msg := "Hello"
	say(&msg)
	println(msg) //변경된 메시지 출력

	says("hi", "yo", "sup")
}

func say(msg *string) {
	println(msg)
	println(*msg)
	*msg = "Changed" //메시지 변경
}

func says(msg... string) {
	for _, m := range msg {
		fmt.Println(m)
	}
}

func sum(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
