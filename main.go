package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/madruga665/gostudies/entities"
	"github.com/madruga665/gostudies/greetings"
	"github.com/madruga665/gostudies/mathOperations"
)

func main() {
	year := time.Now().Year()
	luciano := entities.NewPerson("Luciano", year-1986)
	vanessa := entities.NewPerson("Vanessa", year-1992)
	dante := entities.NewPerson("Dante", year-2021)
	names := []string{"Gladys", "Samantha", "Darrin"}
	message, error := greetings.Hello(luciano.Name)
	messages, error := greetings.Hellos(names)

	sum := mathOperations.Sum(6, 6)
	subtract := mathOperations.Subtract(sum, 2)
	multiply := mathOperations.Multiply(sum, 2)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(message)
	fmt.Println(messages["Gladys"])
	fmt.Println(messages["Samantha"])
	fmt.Println(messages["Darrin"])
	fmt.Println("Oi " + dante.Name + "! esse ano você fez: " + strconv.Itoa(dante.Age) + " anos")
	fmt.Println("Oi " + vanessa.Name + "! esse ano você fez: " + strconv.Itoa(vanessa.Age) + " anos")
	fmt.Println(sum)
	fmt.Println(subtract)
	fmt.Println(multiply)
}
