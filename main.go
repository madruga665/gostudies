package main

import (
	"fmt"
	"log"
	"time"

	"github.com/madruga665/gostudies/entities"
	"github.com/madruga665/gostudies/greetings"
	"github.com/madruga665/gostudies/mathOperations"
)

func main() {
	year := time.Now().Year()
	names := []string{"Gladys", "Samantha", "Darrin"}
	messages, hellosError := greetings.Hellos(names)
	sum := mathOperations.Sum(6, 6)
	subtract := mathOperations.Subtract(sum, 2)
	multiply := mathOperations.Multiply(sum, 2)

	if hellosError != nil {
		log.Fatal(hellosError)
	}

	people := []entities.Person{
		entities.NewPerson("Luciano", year-1986),
		entities.NewPerson("Vanessa", year-1992),
		entities.NewPerson("Dante", year-2021),
	}

	for index, person := range people {
		fmt.Printf("%d - Oi %s! esse ano você fez: %d anos \n", index, person.Name, person.Age)
	}

	fmt.Println(messages["Gladys"])
	fmt.Println(messages["Samantha"])
	fmt.Println(messages["Darrin"])
	fmt.Println(sum)
	fmt.Println(subtract)
	fmt.Println(multiply)
}
