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

	responses := make(chan *MyResponse, count)

	for i := 0; i < count; i++ {
		go DoGet("https://golang.org", &done, responses)
	}

	// will wait till all the goroutines notify Done
	done.Wait()
	close(responses)

	for response := range responses {
		fmt.Println(response)
	}
}

type MyResponse struct {
	Body  string
	Error error
}

func DoGet(url string, done *sync.WaitGroup, reponseChannel chan<- *MyResponse) {
	defer done.Done()

	my := new(MyResponse)

	resp, err := http.Get(url)
	if err != nil {
		my.Error = err
		reponseChannel <- my
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	my.Error = err
	my.Body = string(body)
	reponseChannel <- my
}

// STOPMAIN1 OMIT
