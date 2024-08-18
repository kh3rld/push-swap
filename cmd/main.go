package main

import (
	"fmt"
	"os"
	sort "swap/sortStack"
	valid "swap/validator"
)

func main() {
	arg := os.Args[1]
	if valid.Validate(arg) {
		fmt.Println(arg)
		for _, i := range arg {
			k := int(i)
			data := sort.Stack{Data: []int{k}}
			data.Push(1)
		}
	} else {
		fmt.Println("Error")
	}
}
