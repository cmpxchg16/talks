// +build OMIT

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func DoGet(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	return string(body)
}

func GetWebResults(cancel chan bool) (results []string) {
	// START OMIT
	c := make(chan string, 3)

	go func() { c <- DoGet("https://golang.org") }()
	go func() { c <- DoGet("https://tour.golang.org") }()
	go func() { c <- DoGet("https://blog.golang.org") }()

	timeout := time.After(100 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select { // HL
		case result := <-c: // HL
			results = append(results, result)
		case <-timeout: // HL
			fmt.Println("timed out")
			return
		case <-cancel: // HL
			fmt.Println("canceled")
			return
		}
	}
	// STOP OMIT
	return
}

func main() {
	cancel := make(chan bool) // just to simulate case we want to cancel
	results := GetWebResults(cancel)
	fmt.Println(results)
}
