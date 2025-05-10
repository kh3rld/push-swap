package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack operations
func sa(a *[]int) {
	if len(*a) >= 2 {
		(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
	}
}

func sb(b *[]int) {
	if len(*b) >= 2 {
		(*b)[0], (*b)[1] = (*b)[1], (*b)[0]
	}
}

func ss(a, b *[]int) {
	sa(a)
	sb(b)
}

func pa(a, b *[]int) {
	if len(*b) > 0 {
		*a = append([]int{(*b)[0]}, *a...)
		*b = (*b)[1:]
	}
}

func pb(a, b *[]int) {
	if len(*a) > 0 {
		*b = append([]int{(*a)[0]}, *b...)
		*a = (*a)[1:]
	}
}

func ra(a *[]int) {
	if len(*a) > 1 {
		*a = append((*a)[1:], (*a)[0])
	}
}

func rb(b *[]int) {
	if len(*b) > 1 {
		*b = append((*b)[1:], (*b)[0])
	}
}

func rr(a, b *[]int) {
	ra(a)
	rb(b)
}

func rra(a *[]int) {
	if len(*a) > 1 {
		*a = append([]int{(*a)[len(*a)-1]}, (*a)[:len(*a)-1]...)
	}
}

func rrb(b *[]int) {
	if len(*b) > 1 {
		*b = append([]int{(*b)[len(*b)-1]}, (*b)[:len(*b)-1]...)
	}
}

func rrr(a, b *[]int) {
	rra(a)
	rrb(b)
}

// Check if a slice is sorted in ascending order
func isSorted(a []int) bool {
	if len(a) == 0 {
		return true
	}
	for i := 0; i < len(a)-1; i++ {
		if i+1 >= len(a) {
			return false
		}
		if a[i] > a[i+1] {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		return // No arguments: display nothing
	}

	// Parse input stack
	input := strings.Fields(os.Args[1])
	stackA := make([]int, len(input))
	stackB := []int{}
	seen := make(map[int]bool)

	for i, s := range input {
		val, err := strconv.Atoi(s)
		if err != nil || seen[val] {
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
		seen[val] = true
		stackA[i] = val
	}

	// Read instructions from stdin
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		inst := scanner.Text()
		switch inst {
		case "sa":
			sa(&stackA)
		case "sb":
			sb(&stackB)
		case "ss":
			ss(&stackA, &stackB)
		case "pa":
			pa(&stackA, &stackB)
		case "pb":
			pb(&stackA, &stackB)
		case "ra":
			ra(&stackA)
		case "rb":
			rb(&stackB)
		case "rr":
			rr(&stackA, &stackB)
		case "rra":
			rra(&stackA)
		case "rrb":
			rrb(&stackB)
		case "rrr":
			rrr(&stackA, &stackB)
		default:
			fmt.Fprintln(os.Stderr, "Error")
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		return
	}

	if isSorted(stackA) && len(stackB) == 0 {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
