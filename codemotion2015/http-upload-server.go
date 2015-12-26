package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"sync/atomic"
)

// STARTMAIN1 OMIT
// global shared counter, used by all goroutines, // HL
// should synchronized - using atomic.Add and not a mutex // HL
var x int64

// request handler - each handler run in a dedicated goroutine // HL
func upload(w http.ResponseWriter, r *http.Request) {
	out, err := os.Create("/tmp/" + strconv.FormatInt(atomic.AddInt64(&x, 1), 10))

	if err != nil {
		fmt.Println(err)
	}

	// synchronus heaven - get rid of my back caLLbACK HeLL!O_o // HL
	io.Copy(out, r.Body)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	http.HandleFunc("/upload", upload)
	http.ListenAndServe(":8080", nil)
}

// STOPMAIN1 OMIT
