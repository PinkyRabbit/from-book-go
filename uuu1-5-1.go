package main

import (
	"os"
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	for _, urls:= range os.Args[1:] {
		resp, err:= http.Get(urls)
		if err!=nil {
			fmt.Fprintf(os.Stderr,"Ошибка при получении %v\n", err)
			os.Exit(1)
		}else{
			b, err:= ioutil.ReadAll(resp.Body)
			if err!=nil {
				fmt.Fprintf(os.Stderr,"Ошибка чтения тела ответа %v\n", err)
				os.Exit(1)
			}else{
				fmt.Printf("---------\n%s\n",b)
			}
		}
	}
}
