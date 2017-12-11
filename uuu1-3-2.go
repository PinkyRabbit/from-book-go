package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counts:= make(map[string]int)
	var c int
	for _, filename:= range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err!=nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
		} else {
			for _,line:= range strings.Split(string(data), "\n") {
				counts[line] = c
				c++
			}
		}
	}
	// show result
	for index, line:= range counts {
		fmt.Printf("Line %d\t| %s\n", line, index)
	}
}
