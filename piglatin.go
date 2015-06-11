package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	arg := os.Args[1]
	prefix := regexp.MustCompile("([aeiou].*)").Split(arg, 2)
	suffix := regexp.MustCompile("([aeiou].*)").FindString(arg)

	if len(arg) == len(suffix) {
		fmt.Println(arg + "way")
	} else {
		fmt.Println(suffix + prefix[0] + "ay")
	}
}
