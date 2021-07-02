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

type Cow struct {
	Animal
}

type Bird struct {
	Animal
}

type Snake struct {
	Animal
}

func (a Cow) Name() string {
	return a.name
}

func (a Cow) Eat() string {
	return a.food
}

// NOTE: this works but doesn't meet the definition of the Problem
// func (a Animal) Name() string {
// 	return a.name
// }

// func (a Animal) Eat() string {
// 	return a.food
// }

// func (a Animal) Move() string {
// 	return a.locomotion
// }

// func (a Animal) Speak() string {
// 	return a.noise
// }

// NOTE: this works but seems overkill (seen other solutions that printed here)
func (a Cow) Move() string {
	return a.locomotion
}

func (a Cow) Speak() string {
	return a.noise
}

func (a Bird) Name() string {
	return a.name
}

func (a Bird) Eat() string {
	return a.food
}

func (a Bird) Move() string {
	return a.locomotion
}

func (a Bird) Speak() string {
	return a.noise
}

func (a Snake) Name() string {
	return a.name
}

func (a Snake) Eat() string {
	return a.food
}

func (a Snake) Move() string {
	return a.locomotion
}

func (a Snake) Speak() string {
	return a.noise
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
	fmt.Printf("entries: %v\n", entries)
}

func processLine(entries *[]interface{ Entry }, line string) {
	words := strings.Fields(line)
	if len(words) == 3 {
		result := ""
		commandStr := words[0]
		nameStr := words[1]
		paramStr := words[2]
		switch commandStr {
		case "newanimal":
			*entries = append(*entries, buildEntry(nameStr, paramStr))
			//printEntryList(entries)
		case "query":
			for _, e := range *entries {
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

func testLines(entries *[]interface{ Entry }) {
	processLine(entries, "newanimal aaa cow")
	printEntryList(*entries)
	processLine(entries, "query aaa eat")
	processLine(entries, "query aaa move")
	processLine(entries, "query aaa speak")

	processLine(entries, "newanimal bbb bird")
	printEntryList(*entries)
	processLine(entries, "query bbb eat")
	processLine(entries, "query bbb move")
	processLine(entries, "query bbb speak")

	processLine(entries, "newanimal ccc snake")
	printEntryList(*entries)
	processLine(entries, "query ccc eat")
	processLine(entries, "query ccc move")
	processLine(entries, "query ccc speak")

}

func main() {
	var entries []interface{ Entry }
	for true {
		// prompt and input as string
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		if strings.TrimSpace(line) == "test" {
			testLines(&entries)
		}
		processLine(&entries, line)
	}
}
