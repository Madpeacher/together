package main

import "fmt"

func main() {
	add := func(a, b int) int {
		return a + b
	}

	substract := func(a, b int) int {
		return a - b
	}

	fmt.Println(calculate(1, 2, add))
	fmt.Println(calculate(1, 2, substract))
}

type Operation func(a, b int) int

func calculate(a, b int, operation Operation) int {
	return operation(a, b)
}
