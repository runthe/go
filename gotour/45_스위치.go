// 45_스위치
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")

	switch os := runtime.GOOS; os{
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		//windows, plan9, openbsd
		fmt.Printf("%s", os);
	}
}

