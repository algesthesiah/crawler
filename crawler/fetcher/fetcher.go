package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"time"
)

var rateLimiter = time.Tick(10 * time.Millisecond)

func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	readerBody := bufio.NewReader(response.Body)
	enCode := determineEncoding(readerBody)
	utf8Reader := transform.NewReader(readerBody, enCode.NewDecoder())
	if response.StatusCode == http.StatusOK {
		return ioutil.ReadAll(utf8Reader)
		//if e != nil {
		//	panic(e)
		//}
		//printCityList(all)
		//fmt.Printf("%s\n", all)
	} else {
		return nil, fmt.Errorf("出错了:%d", response.StatusCode)
	}
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, e := r.Peek(1024)
	if e != nil {
		return unicode.UTF8
	}
	enCode, _, _ := charset.DetermineEncoding(bytes, "")
	return enCode
}
