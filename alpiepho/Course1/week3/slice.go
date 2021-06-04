package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// initialize slice of integers
	values := make([]int, 0, 3)

	// loop asking for new values
	var svalue string
	for true {
		// prompt and input as string
		fmt.Print("enter int ('X' to quit): ")
		_, err := fmt.Scanf("%s", &svalue)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// check for exit
		if svalue == "x" || svalue == "X" {
			break
		}
		// parse the input to an integer
		ivalue, err := strconv.Atoi(svalue)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// add new value
		values = append(values, ivalue)

		// print sorted slice
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		fmt.Printf("%v\n", values)
	}
}
