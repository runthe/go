package main

import (
	"fmt"
	"reflect"
)

func main() {
	rects := []struct{ w, h int }{{1, 2}, {2, 3}, {}}

	for _, r := range rects {
		fmt.Printf("(%d,%d)", r.w, r.h)
		fmt.Printf("%#v", r)

	}

	tType := reflect.TypeOf(Item{})

	for i := 0; i < tType.NumField(); i++ {
		fmt.Println(tType.Field(i).Tag)
	}
}

type Item struct {
	name     string "상품 이름"
	price    float64 "상품 가격"
	quantity int "구매 수량"
}