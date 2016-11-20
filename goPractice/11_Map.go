package main

import "fmt"

func main() {

	var idMap map[int]string

	fmt.Println(idMap)

	idMap = make(map[int]string)

	fmt.Println(idMap)

	tickers := map[string]string{
		"GOOG" : "Google inc",
		"MSFT": "Microsoft",
		"FB":"FaceBook",
	}

	fmt.Println(tickers)

	var m map[int]string
	m = make(map[int]string)

	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	noData := m[999]
	fmt.Println(m)
	fmt.Println(noData)

	val, exists := tickers["MSFT"]
	if !exists {
		println("No MSFT ticker")
	}
	fmt.Println(val)

	for key, val := range tickers {
		fmt.Println(key, val)
	}

}