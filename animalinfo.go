package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {
	food       string
	locomotion string
	sound      string
}

type Bird struct {
	food       string
	locomotion string
	sound      string
}

type Snake struct {
	food       string
	locomotion string
	sound      string
}

func CreateAnimal(name, animalType string) Animal {
	switch strings.ToLower(animalType) {
	case "cow":
		return Cow{food: "grass", locomotion: "walk", sound: "moo"}
	case "bird":
		return Bird{food: "worms", locomotion: "fly", sound: "peep"}
	case "snake":
		return Snake{food: "mice", locomotion: "slither", sound: "hsss"}
	default:
		fmt.Println("Invalid animal type. Please enter one of the following types: 'cow', 'bird', or 'snake'")
		return nil
	}
}

func GetAnimalInfo(name, info string, animalMap map[string]Animal) {
	if animal, ok := animalMap[name]; ok {
		if animal != nil {
			switch strings.ToLower(info) {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Please use 'eat', 'move' or 'speak'")
			}
		} else {
			fmt.Println("Animal not found.")
		}
	} else {
		fmt.Println("Animal not found.")
	}
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.sound)
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.sound)
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.sound)
}

func main() {
	animalMap := make(map[string]Animal)
	for {
		fmt.Println(">")
		var userInput string
		reader := bufio.NewReader(os.Stdin)
		userInput, _ = reader.ReadString('\n')
		userInput = strings.ToLower(strings.TrimSpace(userInput))
		words := strings.Fields(userInput)
		if len(words) == 3 {
			if strings.ToLower(words[0]) == "newanimal" {
				newAnimal := CreateAnimal(words[1], words[2])
				if newAnimal != nil {
					animalMap[words[1]] = newAnimal
					fmt.Println("Created it!")
				}
			} else if strings.ToLower(words[0]) == "query" {
				GetAnimalInfo(words[1], words[2], animalMap)
			}
		} else {
			fmt.Println("Please enter 2 strings after the command")
		}
	}
}
