package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}

type Operation func(r rune)

func calculate(a, b int, operation Operation) int {
	switch operation {
	case '+':

	}
}
