package main

import (
	"fmt"
	valid "swap/validator"
)

func main() {
	s := "1234"
	fmt.Println(valid.Validate(s))
}
