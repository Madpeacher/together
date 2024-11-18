package main

import (
	"fmt"
	"time"
)

type kitten struct {
	name      string
	autoritet string
}

func main() {
	kittens := []kitten{
		{"Belyash", "joskii"},
		{"Nigga", "mafia"},
		{"Lei", "doc"},
	}
	for _, kitt := range kittens {
		go Play(kitt)
	}
	time.Sleep(1 * time.Second)
}

func Play(k kitten) {
	fmt.Println(k.name, "with autoritet", k.autoritet, "start to play")
	time.Sleep(1 * time.Second)
	fmt.Println(k.name, "with autoritet", k.autoritet, "finish play")
}
