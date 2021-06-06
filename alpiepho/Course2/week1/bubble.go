package main

import (
	"fmt"
	"strconv"
)

func Swap(values []int, i int) {
	if i >= 0 && (i+1) < len(values) {
		temp := values[i]
		values[i] = values[i+1]
		values[i+1] = temp
	}
}

func BubbleSort(values []int) {
	// outer loop, n..0, work backwards, last element is biggest
	for i := len(values) - 1; i >= 0; i-- {
		// inner loop, push biggest value down
		for j := 0; j < i; j++ {
			if values[j] > values[j+1] {
				Swap(values, j)
			}
		}
	}
}

func main() {
	// initialize slice of integers
	values := make([]int, 0, 10)

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

		// print slice
		fmt.Printf("%v\n", values)
	}

	// TODO call bubble sort and print
	BubbleSort(values)
	fmt.Printf("%v\n", values)
}
