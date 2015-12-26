package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

// STARTMAIN1 OMIT
// Don't Try This at Home ;D // HL
// const count int = 1000000 // HL
const count int = 1000

func main() {

	var done sync.WaitGroup
	done.Add(count)

	for i := 0; i < count; i++ {
		go DoGet("https://golang.org", &done)
	}

	// will wait till all the goroutines notify Done
	done.Wait()
}

func DoGet(url string, done *sync.WaitGroup) {
	defer done.Done()

	resp, err := http.Get(url)
	if err != nil {
		// handle error
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

// STOPMAIN1 OMIT
