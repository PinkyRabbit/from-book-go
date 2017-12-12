// go run uuu1-6.go https://www.ashipka.pw
// go run uuu1-6.go http://prompolit.ru http://hdkinoteatr.com http://ugreshlib.ru http://personalaudio.ru http://starenie.ru http://tempgun.ru http://www.n-kodeks.ru http://gold.co.ua http://ruconf.ru http://www.caduser.ru http://sc170.kharkov.ua http://larevolution.ru http://tenchi.ru http://be-mag.ru http://soft-m.ru http://nokia.bir.ru http://www.arspas.ru http://rrcdetstvo.ru http://www.vishivay.com http://expertmagazine.ru http://foreverlove.ru
package f1_6

import (
	"time"
	"os"
	"net/http"
	"fmt"
	"io"
	"io/ioutil"
)

func main()  {
	start:= time.Now()
	ch:= make(chan string)
	for _,url:= range os.Args[1:] {
		go fetch(url, ch)
	}
	// получение результатов
	for range os.Args[1:]{
		fmt.Println(<-ch)
	}
	fmt.Println("-------------")
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start:= time.Now()
	resp, err:= http.Get(url)
	if err!=nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err:= io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err!=nil {
		ch <- fmt.Sprintf("Ошибка чтения %s:\n%v", url, err)
		return
	}
	secs:= time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}