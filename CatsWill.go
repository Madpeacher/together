package main

import (
	"fmt"
	"time"
)

type kitten struct {
	name      string
	autoritet string
}

type KittErr struct {
	error string
}

func main() {
	kittens := []kitten{
		{"Belyash", "joskii"},
		{"Nigga", "mafia"},
		{"Lei", "doc"},
		{"Invalid", "Loshok"},
	}
	for _, kitt := range kittens {
		go Play(kitt)
	}
	time.Sleep(1 * time.Second)
}

func checkauto(k kitten) KittErr {
	if k.autoritet == "Loshok" {
		return KittErr{"Kiiten cannot be Loshok"}
	}
	return KittErr{}
}
func Play(k kitten) {
	fmt.Println(k.name, "with autoritet", k.autoritet, "start to play")
	time.Sleep(1 * time.Second)
	fmt.Println(k.name, "with autoritet", k.autoritet, "finish play")
}
