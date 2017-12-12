package f1_5

import (
	"os"
	"io"
	"net/http"
	"fmt"
)

func main() {
	for _,url:=range os.Args[1:] {
		resp, err:= http.Get(url)
		if err!=nil {
			fmt.Fprintf(os.Stderr, "Нет соединения с %v", err)
			os.Exit(1)
		}else{
			b, err:= io.Copy(os.Stdout, resp.Body)
			if err!=nil {
				fmt.Fprintf(os.Stderr, "Не могу прочитать body %v", err)
				os.Exit(1)
			}else{
				fmt.Printf("\n--------\n%v\n", b)
			}
		}

	}
}