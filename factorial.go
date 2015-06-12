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

	input, e := strconv.Atoi(os.Args[1])
	if e != nil {
		fmt.Println(e)
	}

	output := (factorial(input))
	outputstr := strconv.Itoa(output)
	fmt.Println(outputstr)
}
