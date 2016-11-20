package main

import "fmt"

func main() {

	var name string
	var category int = 2

	switch category {
	case 1:
		name = "Paper Book"
	case 2:
		name = "eBook"
	case 3, 4:
		name = "Blog"
	default:
		name = "Other"
	}

	println(name)

	var t interface{}
	t = name

	switch t.(type) {
	case int:
		println("int")
	case bool:
		println("bool")
	case string:
		println("string")
	default:
		println("unknown")
	}

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	for sum < 10000 {
		sum *= 2
		//break
	}

	fmt.Println(sum)

	names := []string{"하하", "호호", "키키"}

	for index, name := range names {
		fmt.Println(index, name)
	}
}


