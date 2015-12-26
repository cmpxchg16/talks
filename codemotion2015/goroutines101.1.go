package main

import (
	"fmt"
	"runtime"
	"time"
)

// STARTMAIN1 OMIT
const count int = 5

func even() {
	i := -1
	for {
		i++
		if i%2 == 0 {
			fmt.Printf("even is the dominant (%v)!\n", i)
			if i < count {
				runtime.Gosched()
			}
		}
	}
}

func odd() {
	i := 0
	for {
		i++
		if i%2 == 1 {
			fmt.Printf("odd is the dominant (%v)!\n", i)
			runtime.Gosched()
		}
	}
}

func main() {
	runtime.GOMAXPROCS(1) // only one OS native thread
	even()
	odd()
	time.Sleep(time.Second)
}

// STOPMAIN1 OMIT
