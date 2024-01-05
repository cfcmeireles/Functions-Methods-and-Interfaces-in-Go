package main

import (
	"fmt"
	"strings"
)

type animal struct {
	eat   string
	move  string
	speak string
}

func main() {
	fmt.Println("Enter the animal you'd like to get more information about:")
	var animalPrompt string
	fmt.Scanln(&animalPrompt)
	animalPrompt = strings.ToLower(animalPrompt)
	fmt.Println("Eat, move or speak?")
	var activityPrompt string
	fmt.Scanln(&activityPrompt)
	activityPrompt = strings.ToLower(activityPrompt)

	cow := animal{
		eat:   "grass",
		move:  "walk",
		speak: "moo",
	}

	bird := animal{
		eat:   "worms",
		move:  "fly",
		speak: "peep",
	}

	snake := animal{
		eat:   "mice",
		move:  "slither",
		speak: "hsss",
	}

	if animalPrompt == "cow" && activityPrompt == "eat" {
		fmt.Println(cow.eat)
	} else if animalPrompt == "cow" && activityPrompt == "move" {
		fmt.Println(cow.move)
	} else if animalPrompt == "cow" && activityPrompt == "speak" {
		fmt.Println(cow.speak)
	}

	if animalPrompt == "bird" && activityPrompt == "eat" {
		fmt.Println(bird.eat)
	} else if animalPrompt == "bird" && activityPrompt == "move" {
		fmt.Println(bird.move)
	} else if animalPrompt == "bird" && activityPrompt == "speak" {
		fmt.Println(bird.speak)
	}

	if animalPrompt == "snake" && activityPrompt == "eat" {
		fmt.Println(snake.eat)
	} else if animalPrompt == "snake" && activityPrompt == "move" {
		fmt.Println(snake.move)
	} else if animalPrompt == "snake" && activityPrompt == "speak" {
		fmt.Println(snake.speak)
	}

}
