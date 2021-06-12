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

func buildEntry(name, animal string) interface{ Entry } {
	var result Entry

	switch animal {
	case "cow":
		result = Cow{
			Animal{
				name:       name,
				food:       "grass",
				locomotion: "walk",
				noise:      "moo",
			},
		}
	case "bird":
		result = Bird{
			Animal{
				name:       name,
				food:       "worms",
				locomotion: "fly",
				noise:      "peep",
			},
		}
	case "snake":
		result = Snake{
			Animal{
				name:       name,
				food:       "mice",
				locomotion: "slither",
				noise:      "hiss",
			},
		}
	}
	return result
}

func printEntryList(entries []interface{ Entry }) {
	for _, e := range entries {
		fmt.Println(e.Name())
		fmt.Println(e.Eat())
		fmt.Println(e.Move())
		fmt.Println(e.Speak())
		fmt.Println()
	}
}

func main() {
	var entries []interface{ Entry }
	//entries = append(entries, buildEntry("cow", "cow"))
	//entries = append(entries, buildEntry("bird", "bird"))
	//entries = append(entries, buildEntry("snake", "snake"))
	//printEntryList(entries)

	for true {
		// prompt and input as string
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		// words should be: newanimal, name, animal   or   query, name, trait
		words := strings.Fields(line)
		if len(words) == 3 {
			result := ""
			commandStr := words[0]
			nameStr := words[1]
			paramStr := words[2]
			switch commandStr {
			case "newanimal":
				entries = append(entries, buildEntry(nameStr, paramStr))
				//printEntryList(entries)
			case "query":
				for _, e := range entries {
					if e.Name() == nameStr {
						switch paramStr {
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
}
