package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFail = errors.New("request fail")


func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://bbs.ruliweb.com/best/humor",
		"https://dvdprime.com/g2/bbs/board.php?bo_table=comm&wr_id=21199644",
		"https://www.44bits.io/ko/post/easy-deploy-with-docker",
		"https://gall.dcinside.com/mgallery/board/view/?id=war&no=1700019&page=1",
		"https://naver.com",
		"https://google.com",
		"https://amazon.com",
	}
	results["gello"] = "hello"
	for _, url := range urls {
		result := "OK"
		err := hitURL(url)
		if err != nil {
			result = "fail"
		}
		results[url] = result
	}
	for url, result := range results{
		fmt.Println(url, result)

	}
}


func hitURL(url string) error {
	fmt.Println("checking: ", url)
	res, err := http.Get(url)
	if err != nil || res.StatusCode >=400{
		fmt.Println(err, res.StatusCode)
 		return errRequestFail
	}
	return nil
}