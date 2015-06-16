package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	m := make(map[string]float64)
	m["USD"] = 1.00
	m["EUR"] = 0.89
	m["GBP"] = 1.54

	input, err := strconv.ParseFloat(os.Args[1], 64)

	if err != nil {
		fmt.Println(err)
	}

	diff := input - m[os.Args[3]]

	output := input * (m[os.Args[3]] + diff) / diff

	fmt.Println(output)
}
