package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main() {
	counts:= make(map[string]int)
	var hasDouble bool
	for _, file:= range os.Args[1:] {
		lines, err := ioutil.ReadFile(file)
		if err!=nil {
			fmt.Fprintf(os.Stderr, "pr 1-3-1: %v", err)
		} else {
			hasDouble = false
			for index, line:= range strings.Split(string(lines), "\n"){
				counts[line]++
				if counts[line]>1{
					fmt.Printf("%s line %d\n", file, index)
					fmt.Println(string(line))
					hasDouble = true
				}
			}
			if hasDouble {
				fmt.Println(file)
			}
		}
	}
}