package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

type Cow struct {
	Animal
}

type Bird struct {
	Animal
}

type Snake struct {
	Animal
}

func main() {
	cow := Cow{
		Animal{
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
	}
	bird := Bird{
		Animal{
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
	}
	snake := Snake{
		Animal{
			food:       "mice",
			locomotion: "slither",
			noise:      "hiss",
		},
	}

	for true {
		// prompt and input as string
		fmt.Print(">")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		words := strings.Fields(line) // animal, trait
		if len(words) == 2 {
			result := ""
			animalStr := words[0]
			traitStr := words[1]
			switch animalStr {
			case "cow":
				switch traitStr {
				case "eat":
					result = cow.Eat()
				case "move":
					result = cow.Move()
				case "speak":
					result = cow.Speak()
				}
			case "bird":
				switch traitStr {
				case "eat":
					result = bird.Eat()
				case "move":
					result = bird.Move()
				case "speak":
					result = bird.Speak()
				}
			case "snake":
				switch traitStr {
				case "eat":
					result = snake.Eat()
				case "move":
					result = snake.Move()
				case "speak":
					result = snake.Speak()
				}
			}
			if len(result) > 0 {
				fmt.Println(result)
			}
		}
	}
}
