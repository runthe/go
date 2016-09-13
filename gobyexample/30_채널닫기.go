/*
    author twitter : @mmcgrana
	https://gobyexample.com/closing-channels
    Go by Example: closing-channels

	채널을 닫는 것은 더이상 어떤 값도 그것으로 보내지지 않을 것이란 것을 가리킨다.
	이것은 채널의 리시버와 통신은 완료하는데 유용하다.
	여기서 worker는 따로 일하는 고루틴을 말하는 듯하다(역주)
*/

package main

import "fmt"

// 이 예제에서, 우리는 jobs 채널을 이용하여 main() 고루틴으로부터 work고루틴으로 작업이 완료되었다고
// 통신할 것이다. worker로 할일이 더 이상 없다면 우리는 jobs채널을 닫을 것이다.

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("send job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}
