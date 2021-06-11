package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Entry interface {
	Name() string
	Eat() string
	Move() string
	Speak() string
}

type Animal struct {
	name       string
	food       string
	locomotion string
	noise      string
}

func (a Animal) Name() string {
	return a.name
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

func printEntryList(entries []interface{ Entry }) {
	for _, e := range entries {
		fmt.Println(e.Name())
		fmt.Println(e.Eat())
		fmt.Println(e.Move())
		fmt.Println(e.Speak())
	}
}

func main() {
	cow := Cow{
		Animal{
			name:       "cow",
			food:       "grass",
			locomotion: "walk",
			noise:      "moo",
		},
	}
	bird := Bird{
		Animal{
			name:       "bird",
			food:       "worms",
			locomotion: "fly",
			noise:      "peep",
		},
	}
	snake := Snake{
		Animal{
			name:       "snake",
			food:       "mice",
			locomotion: "slither",
			noise:      "hiss",
		},
	}
	var entries []interface{ Entry }
	entries = append(entries, cow)
	entries = append(entries, bird)
	entries = append(entries, snake)
	printEntryList(entries)

	for true {
		// prompt and input as string
		fmt.Print("> ")
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

			for _, e := range entries {
				if e.Name() == animalStr {
					switch traitStr {
					case "eat":
						result = e.Eat()
					case "move":
						result = e.Move()
					case "speak":
						result = e.Speak()
					}
				}
			}
			if len(result) > 0 {
				fmt.Println(result)
			}
		}
	}
}
