package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter float: ")
	var fvalue float64
	_, err := fmt.Scanf("%f", &fvalue)
	if err != nil {
		fmt.Println(err)
	}
	ivalue := int(fvalue)
	fmt.Println(ivalue)

}
