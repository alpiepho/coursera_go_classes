package main

import (
	"fmt"
	"strconv"
)

func GenDisplaceFn(a, vo, so float64) func(t float64) float64 {
	return func(t float64) float64 {
		// s = ½ a t^2 + vo*t + so
		return (0.5 * a * t * t) + (vo * t) + so
	}
}

func getFlaot64Input(prompt string) (float64, error) {
	var line string
	var result float64

	fmt.Printf("%s", prompt)
	_, err := fmt.Scanf("%s", &line)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	result, err = strconv.ParseFloat(line, 64)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	return result, nil
}

func main() {
	var err error
	var a float64
	var vo float64
	var so float64
	var t float64

	a, err = getFlaot64Input("acceleration: ")
	if err != nil {
		return
	}
	vo, err = getFlaot64Input("initial velocity: ")
	if err != nil {
		return
	}
	so, err = getFlaot64Input("initial displacement: ")
	if err != nil {
		return
	}
	t, err = getFlaot64Input("time: ")
	if err != nil {
		return
	}

	fn := GenDisplaceFn(a, vo, so)
	fmt.Printf("fn(%v):\n", t)
	fmt.Println(fn(t))
}
