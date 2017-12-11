package f1_3

import (
	"bufio"
	"fmt"
	"os"
	//"os/exec"
)

var index int

func main() {
	/* dup 1
	counts := make(map[string]int)
	var c int
	input := bufio.NewScanner(os.Stdin)
	for len(counts)<4 && input.Scan()  {
		counts[input.Text()] = c
		c++
		//fmt.Println(input.Text())
	}
	fmt.Println(counts)
	for line, index:= range counts  {
		fmt.Printf("Output %d line = %s\n", index, line)
	}
	fmt.Println("=========")
	*/
	counts:= make(map[string]int)
	//var t int
	files:= os.Args[1:]
	for _, fileName:= range files {
		countLines(fileName, counts)
	}

	if(len(counts)==0){
		var t int
		input:= bufio.NewScanner(os.Stdin)
		for len(counts)==0 && t<5 && input.Scan() {
			if t==5 {
				fmt.Println("Последняя попытка найти файл!")
			}
			countLines(input.Text(), counts)
			t++
		}
	}

	for line, i:= range counts {
		fmt.Printf("Line %d\t| %s\n", i, line)
	}
}

func countLines(fileName string, counts map[string]int) {
	f, err:= os.Open(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения ", fileName)
	}else{
		input:= bufio.NewScanner(f)
		for input.Scan() {
			counts[input.Text()] = index
			index ++
		}
		//fmt.Println(counts)
		f.Close()
		//fmt.Println("Прочитано ", counts)
	}
}