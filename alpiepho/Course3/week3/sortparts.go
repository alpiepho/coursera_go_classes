package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func sortPart1(values []int, index1, index2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("sortPart1 before: %v\n", values[index1:index2])
	sort.Slice(values, func(i, j int) bool {
		if i >= index1 && i < index2 && j >= index1 && j < index2 {
			return values[i] < values[j]
		}
		return false
	})
	fmt.Printf("sortPart1 after:  %v\n", values[index1:index2])
}

func sortPart2(values []int, index1, index2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("sortPart2 before: %v\n", values[index1:index2])
	sort.Slice(values, func(i, j int) bool {
		if i >= index1 && i < index2 && j >= index1 && j < index2 {
			return values[i] < values[j]
		}
		return false
	})
	fmt.Printf("sortPart2 after:  %v\n", values[index1:index2])
}

func sortPart3(values []int, index1, index2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("sortPart3 before: %v\n", values[index1:index2])
	sort.Slice(values, func(i, j int) bool {
		if i >= index1 && i < index2 && j >= index1 && j < index2 {
			return values[i] < values[j]
		}
		return false
	})
	fmt.Printf("sortPart3 after:  %v\n", values[index1:index2])
}

func sortPart4(values []int, index1, index2 int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("sortPart4 before: %v\n", values[index1:index2])
	sort.Slice(values, func(i, j int) bool {
		if i >= index1 && i < index2 && j >= index1 && j < index2 {
			return values[i] < values[j]
		}
		return false
	})
	fmt.Printf("sortPart4 after:  %v\n", values[index1:index2])
}

func sortFinal(values []int) {
	var wg sync.WaitGroup

	partLen := len(values) / 4
	index := 0
	wg.Add(1)
	go sortPart1(values, index, index+partLen, &wg)
	index = index + partLen
	wg.Add(1)
	go sortPart2(values, index, index+partLen, &wg)
	index = index + partLen
	wg.Add(1)
	go sortPart3(values, index, index+partLen, &wg)
	index = index + partLen
	wg.Add(1)
	go sortPart4(values, index, len(values), &wg)
	index = index + partLen
	wg.Wait()

	// final sort values
	fmt.Printf("final     before: %v\n", values)
	sort.Slice(values, func(i, j int) bool {
		return values[i] < values[j]
	})

	fmt.Printf("final     after:  %v\n", values)
}

func test() {
	fmt.Println()
	var values1 = []int{12, 11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sortFinal(values1)
	fmt.Println()
	var values2 = []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sortFinal(values2)
	fmt.Println()
	var values3 = []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	sortFinal(values3)
	fmt.Println()
	var values4 = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	sortFinal(values4)
	fmt.Println()
	var values5 = []int{8, 7, 6, 5, 4, 3, 2, 1}
	sortFinal(values5)
	fmt.Println()
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
		// check for test
		if svalue == "test" {
			test()
			return
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

	// run the sort algoritm
	sortFinal(values)
}
