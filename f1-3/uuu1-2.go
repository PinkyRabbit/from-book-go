package f1_3

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)

	str, sep2 := "", ""
	for _, arg := range os.Args[1:] {
		str += sep2 + arg
		sep2 = " "
	}
	fmt.Println(str)
	fmt.Println("-------")
	fmt.Println(strings.Join(os.Args[1:], " "))
	fmt.Println(os.Args[1:])

	fmt.Println("==========")
	fmt.Println("1 :")
	fmt.Println(os.Args)
	fmt.Println("2 :")
	for j := 0; j < len(os.Args); j++ {
		fmt.Printf("os.Args[%d] = "+os.Args[j]+"\n", j)
	}
	fmt.Println("3 :")

}
