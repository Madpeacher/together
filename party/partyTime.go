package main

import (
	"fmt"
	"strconv"
)

type phones struct {
	name_phone string
	model      int
}

type Guest struct {
	name       string
	age        int
	city       string
	chromosome int
	phones
}

type Chill interface {
	getChill()
}

type DJ struct {
	name     string
	Playlist []string
}

type Narco struct {
	Chromosome int
	NarcoPrice float64
}

func main() {
	Rasul := Guest{
		name:       "Rasul",
		age:        18,
		city:       "Buinaksk",
		chromosome: 46,
		phones:     phones{"Xiaomi", 5},
	}

	Jabrik := Guest{
		name:       "Jabrik",
		age:        27,
		city:       "Buinaksk",
		chromosome: 45,
		phones:     phones{"Iphone", 6},
	}

	Alena := Guest{
		name:       "Alena",
		age:        17,
		city:       "Cursk",
		chromosome: 46,
		phones:     phones{"Xiaomi", 1},
	}
	Tom := DJ{
		name:     "Tom",
		Playlist: []string{"imagine dragons", "bad omens"},
	}

	Alex := Narco{
		Chromosome: 32,
		NarcoPrice: 5000}

	Alena.Changechromosome()

	//Rasul.Changechromosome()
	Rasul.Meeting(&Alena)
	Rasul.Meeting(&Jabrik)
	printGuest(Rasul)
	printGuest(Jabrik)
	printGuest(Alena)

	fmt.Println(Alex)
	fmt.Println(Tom)

	evrybodyHandsUp(&Alex)
	evrybodyHandsUp(&Jabrik)

	Alena.giveChromosome(Rasul)

}

func evrybodyHandsUp(chillman Chill) {
	fmt.Println("evrybodyHandsUp")
	guest, ok := chillman.(*Guest)
	if ok {
		fmt.Println("поприветствуем нашего гостяя - ", guest.name)
	}
}

func (*Narco) getChill() {
	fmt.Println("Chilll но с опаской")
}

func (*Guest) getChill() {
	fmt.Println("Chilllllll broooo")
}

func (guest1 *Guest) giveChromosome(guest2 Guest) {
	guest1.chromosome--
	guest2.chromosome++
}

func (chelick *Guest) Changechromosome() {
	chelick.chromosome++
}

func (meet *Guest) Meeting(g *Guest) {
	if meet.chromosome == g.chromosome {
		fmt.Println("ДА КАК ВЫ НИПАНИМАЕТЕ ОНИ ЯБУЦАЦАЦААЦАЦЦААЦА")
	} else {
		fmt.Println(g.name, "МЫ С ТОБОЮ НЕ ПАРА НЕ ПАРА ВОТ ТАКАЯ У НАС  СТОБОЮ ЗАПАРА ЗАПАРА")
	}
}

func printGuest(guest Guest) {
	fmt.Println("Имя: " + guest.name)
	fmt.Println("Лет: " + strconv.Itoa(guest.age))
	fmt.Println("хромосом: " + strconv.Itoa(guest.chromosome))
	fmt.Println("с телефоном: " + guest.name_phone)
	fmt.Println()
}
