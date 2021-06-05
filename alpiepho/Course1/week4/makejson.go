package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// initialize map
	values := make(map[string]string)

	// prompt and save to map
	var svalue string
	fmt.Print("enter name: ")
	_, err := fmt.Scanf("%s", &svalue)
	if err != nil {
		fmt.Println(err)
		return
	}
	values["name"] = svalue
	fmt.Print("enter address: ")
	_, err = fmt.Scanf("%s", &svalue)
	if err != nil {
		fmt.Println(err)
		return
	}
	values["address"] = svalue

	// convert to json and print
	asJson, _ := json.Marshal(values)
	fmt.Println(string(asJson))
}
