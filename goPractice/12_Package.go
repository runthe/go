package main

import "fmt"

var pop map[string]string

func init() {
	fmt.Println("test")
	pop = make(map[string]string)
}

func main() {
	fmt.Println("main")
}