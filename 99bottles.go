package main

import (
	"bytes"
	"fmt"
	"strconv"
)

func main() {
	var buffer bytes.Buffer

	for i := 99; i > 0; i-- {
		if i == 1 {
			buffer.WriteString(strconv.Itoa(i) + " bottle of beer on the wall\nTake one down, pass it around...")
		} else {
			buffer.WriteString(strconv.Itoa(i) + " bottles of beer on the wall\nTake one down, pass it around...\n")
		}
	}

	fmt.Println(buffer.String())
	fmt.Println("Now there's no more bottles of beer on the wall!")
}
