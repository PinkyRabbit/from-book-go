package main

import (
	"os"
	"strings"
	"net/http"
	"fmt"
	"io"
)

func main() {
	for _, urls:= range os.Args[1:] {
		url:= urls
		if strings.HasPrefix(url, "http://")==false {
			url= "http://"+url
		}
		resp, err:= http.Get(url)
		if err!=nil {
			fmt.Fprintf(os.Stderr, "Ошибка подключения к ссылке %v", err)
			os.Exit(1)
		}else{
			b, err:= io.Copy(os.Stdout, resp.Body)
			if err!=nil {
				fmt.Fprintf(os.Stderr, "Ошибка чтения тела ответа %v", err)
				os.Exit(1)
			}else{
				fmt.Printf("\n-------\n%v\n", b)
			}
		}
	}
}
