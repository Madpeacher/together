package main

import (
	"fmt"
	"time"
)

func Catch(mice chan string) {
	for i := 0; i < 5; i++ {
		mouse := fmt.Sprintf("mice %d", i)
		fmt.Println("Cathes mouse", mouse)
		mice <- mouse
		time.Sleep(1 * time.Second)
	}
	close(mice)
}

func Eat(mice chan string) {
	for mpuse := range mice {
		fmt.Println("Eat mpuse", mpuse)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	mice := make(chan string)
	go Catch(mice)
	go Eat(mice)

	time.Sleep(3 * time.Second)
}
