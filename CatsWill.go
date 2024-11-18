package main

import (
	"errors"
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
		{"Invalid", "Loshok"},
	}
	for _, kitt := range kittens {
		go Play(kitt)
	}
	time.Sleep(1 * time.Second)
	fmt.Println(checkauto(kittens[3]))
}

type KittErr struct {
	err error
}

func (k *KittErr) Error() string {
	return k.err.Error()
}

func exp() error {
	return &KittErr{
		err: errors.New("kitten cannot be loshok"),
	}
}
func checkauto(k kitten) error {
	if k.autoritet == "Loshok" {
		return exp()
	}
	return nil
}
func Play(k kitten) {
	fmt.Println(k.name, "with autoritet", k.autoritet, "start to play")
	time.Sleep(1 * time.Second)
	fmt.Println(k.name, "with autoritet", k.autoritet, "finish play")
}
