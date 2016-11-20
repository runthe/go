package main

import "fmt"

func main() {

	var a int
	fmt.Println(a)

	var f float32
	fmt.Println(f)

	var b, c, d int = 1, 2, 3
	fmt.Println(b, c, d)

	var e = 1
	var g = "gg"

	fmt.Println(e, g)

	const h int = 1
	const i string = "hi"

	const (
		j = "j"
		k = "k"
		l = "l"
	)

	const (
		Apple = iota
		Grape
		Orage
	)

	fmt.Println(Apple, Grape)

	rawLiteral := `니가 먼데
	ㅇㅅㅇ \n ㅋㅋ`
	interLiteral := "하하 먼데 \n ㅋㅋ"

	fmt.Println(rawLiteral, interLiteral)

	fmt.Println(string([]byte("크크크")))
	fmt.Println(string([]rune("크크크")))
}
