package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Enter string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	svalue := scanner.Text()

	// look for:
	// starts with 'i' or 'I'
	// ends with 'n' or 'N'
	// contains 'a' or 'A'
	var found bool = false
	var length int = len(svalue)
	if length >= 3 {
		if svalue[0] == 'i' || svalue[0] == 'I' {
			if svalue[length-1] == 'n' || svalue[length-1] == 'N' {
				for i := 1; i < length-1; i++ {
					if svalue[i] == 'a' || svalue[i] == 'A' {
						found = true
						break
					}
				}
			}
		}
	}
	if found {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
