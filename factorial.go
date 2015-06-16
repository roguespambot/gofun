package main

import (
	"fmt"
	"os"
	"strconv"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {
	input, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}

	output := (factorial(input))
	outputstr := strconv.Itoa(output)
	fmt.Println(outputstr)
}
