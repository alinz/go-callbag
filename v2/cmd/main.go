package main

import (
	"fmt"

	callbag "github.com/alinz/go-callbag/v2"
)

type Value = callbag.Value

func main() {
	fromSlice := callbag.FromValues(1, 2, 3)
	forEach := callbag.ForEach(func(val Value) {
		fmt.Println(val)
	})

	forEach.Call(fromSlice)
}
