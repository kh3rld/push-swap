package main

import (
	"fmt"
	"os"
	"strings"

	pushswap "swap/validator"
)

func main() {
	args := os.Args[1:]
	values, err := pushswap.ParseArgs(args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}
	if values == nil {
		fmt.Println()
		return
	}

	stack := pushswap.NewStackList()
	for _, val := range values {
		stack.Push(val)
	}

	if !stack.IsSorted() {
		switch stack.Length() {
		case 2:
			pushswap.SwitchFirstTwo(stack, "a")
		case 3:
			pushswap.TinySort(stack)
		default:
			pushswap.Sort_stack(stack)
		}
	}

	fmt.Println(strings.Join(pushswap.Moves, "\n"))
}
