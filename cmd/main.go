package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"swap/sortStack"
	"swap/validator"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]

	if !validator.Validate(arg) {
		fmt.Println("Error")
		return
	}

	values := strings.Split(arg, " ")
	a := sortStack.NewStack()
	b := sortStack.NewStack()

	for _, v := range values {
		data, err := strconv.Atoi(v)
		if err != nil {
			fmt.Println("Error")
			return
		}
		a.Push(data)
	}

	sortStack.SortS(a, b)

	for len(a.Data) > 0 {
		value, _ := a.Pop()
		fmt.Println(value)
	}
}
