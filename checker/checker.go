package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"swap/sortStack"
	"swap/validator"
)

func executeInstructions(a, b *sortStack.Stack, instructions []string) bool {
	for _, instruction := range instructions {
		switch instruction {
		case "pa":
			sortStack.Pa(a, b)
		case "pb":
			sortStack.Pb(a, b)
		case "sa":
			sortStack.Sa(a)
		case "sb":
			sortStack.Sb(b)
		case "ss":
			sortStack.Ss(a, b)
		case "ra":
			sortStack.Ra(a)
		case "rb":
			sortStack.Rb(b)
		case "rr":
			sortStack.Rr(a, b)
		case "rra":
			sortStack.Rra(a)
		case "rrb":
			sortStack.Rrb(b)
		case "rrr":
			sortStack.Rrr(a, b)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return false
		}
	}
	return sortStack.IsSorted(a) && len(b.Data) == 0
}

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]

	if !validator.Validate(arg) {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	values := strings.Split(arg, " ")
	a := sortStack.NewStack()
	b := sortStack.NewStack()

	for _, v := range values {
		data, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		a.Push(data)
	}

	scanner := bufio.NewScanner(os.Stdin)
	var instructions []string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		instructions = append(instructions, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if executeInstructions(a, b, instructions) {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
