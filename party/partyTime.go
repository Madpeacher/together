package main

import (
	"fmt"
	"strconv"
)

type phones struct {
	name  string
	model int
}

type guest struct {
	name       string
	age        int
	city       string
	chromosome int
	phone      *phones
}

func main() {
	Rasul := guest{
		name:       "Rasul",
		age:        18,
		city:       "Buinaksk",
		chromosome: 46,
		phone:      &phones{"Xiaomi", 5},
	}

	Jabrik := guest{
		name:       "Jabrik",
		age:        27,
		city:       "Buinaksk",
		chromosome: 45,
		phone:      &phones{"Iphone", 6},
	}

	Alena := guest{
		name:       "Alena",
		age:        17,
		city:       "Cursk",
		chromosome: 46,
		phone:      &phones{"Xiaomi", 1},
	}
	changeChromosome(&Rasul)

	printGuest(Rasul)
	printGuest(Jabrik)
	printGuest(Alena)
}

func changeChromosome(chelick *guest) {
	chelick.chromosome++
}

func printGuest(guest guest) {
	fmt.Println("Имя: " + guest.name)
	fmt.Println("Лет: " + strconv.Itoa(guest.age))
	fmt.Println("хромосом: " + strconv.Itoa(guest.chromosome))
	fmt.Println("с телефоном: " + guest.phone.name)
	fmt.Println()
}
