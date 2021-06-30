package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

type Person struct {
	fname string
	lname string
}

func main() {

	names := list.New()
	scanner := bufio.NewScanner(os.Stdin)

	// read filename from user
	var filename string
	fmt.Print("enter filename: ")
	scanner.Scan()
	filename = scanner.Text()

	// attempt to open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s\n", filename)
		return
	}
	// close file when main exits
	defer file.Close()

	// read each line of the file and add struct to slice
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			var person Person
			person = Person{fname: parts[0], lname: parts[1]}
			names.PushBack(person)
		} else {
			fmt.Printf("Error reading file, bad line: %s\n", line)
		}
	}

	// check for reading errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s\n", filename)
		return
	}

	// cycle thru slice and print first/last names
	for e := names.Front(); e != nil; e = e.Next() {
		person := Person(e.Value.(Person))
		fmt.Printf("%s %s\n", person.fname, person.lname)
	}
}
