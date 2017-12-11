package main

import (
	"os"
	"net/http"
	"fmt"
)

func main()  {
	for _, urls:=range os.Args[1:] {
		resp, err:= http.Get(urls)
		if err!=nil {
			fmt.Fprintf(os.Stderr, "Ошибка подключения к сайту %v", err)
			os.Exit(1)
		}else{
			fmt.Println(resp.StatusCode)
		}
	}
}
