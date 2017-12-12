package f1_6

import (
	"time"
	"fmt"
	"os"
	"net/http"
	"io"
	"io/ioutil"
)
// for test:
// http://zforum.pl http://nwlegalhouse.ru http://4pera.ru http://a-4.ru http://polariss.ru http://vkirove.ru http://powerbibbs.com http://99hainan.com http://bumprompak.by http://tavernaute.com http://digitaltutor.forumcity.com http://bbs.byr.cc http://contrafakt.com http://0577183.com http://uu88s.com http://mengyem.com http://gong-shi.com http://forum.chuwi.com http://dashenkashop.de http://zsko.ru http://arkh-edu.ru http://caduser.ru http://gold.co.ua http://arspas.ru http://paterton.ru http://play-progress.com http://julianna.ru http://rastrwin.ru http://allprint-service.ru http://kemoko.su

func main() {
	secs:= time.Now()
	ch:= make(chan string)
	for _,url:= range os.Args[1:] {
		go fetchall(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Println("-----")
	fmt.Printf("%.fs elapsed\n", time.Since(secs).Seconds())
}

func fetchall(url string, ch chan<- string) {
	var result1 [3]float64
	var	result2 [3]int64
	for i:=0;i<3;i++ {
		start:= time.Now()
		resp, err:= http.Get(url)
		if err!=nil {
			result1[i] = -1
			result2[i] = 0
		}else{
			result2[i], err = io.Copy(ioutil.Discard, resp.Body)
			resp.Body.Close()
			if err!=nil {
				result2[i] = 0
			}
			result1[i] = time.Since(start).Seconds()
		}
	}
	ch <- fmt.Sprintf("%.2fs\t%7d\t%.2fs\t%7d\t%.2fs\t%7d\t%s", result1[0], result2[0], result1[1], result2[1], result1[2], result2[2], url)
}