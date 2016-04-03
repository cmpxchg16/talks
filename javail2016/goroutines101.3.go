package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"time"
)

// STARTMAIN1 OMIT
func DoGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Printf("url: %v, status code: %v, body len: %v\n", url, resp.StatusCode, len(body))
}

var i int

func f1() {
	for {
		DoGet("https://golang.org")
		fmt.Println(i)
	}
}

func f2() {
	for {
		i++
	}
}

func main() {
	runtime.GOMAXPROCS(1) // only one OS native thread
	go f1()
	time.Sleep(time.Second)
	go f2()
	time.Sleep(3 * time.Second)
}

// STOPMAIN1 OMIT
