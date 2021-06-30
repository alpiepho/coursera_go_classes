package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// initialize map
	values := make(map[string]string)
	scanner := bufio.NewScanner(os.Stdin)

	// prompt and save to map
	var svalue string
	fmt.Print("enter name: ")
	scanner.Scan()
	svalue = scanner.Text()
	values["name"] = svalue
	fmt.Print("enter address: ")
	scanner.Scan()
	svalue = scanner.Text()
	values["address"] = svalue

	// convert to json and print
	asJson, _ := json.Marshal(values)
	fmt.Println(string(asJson))
}
