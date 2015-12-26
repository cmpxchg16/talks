package main

import "runtime"

//obvious race O_o // HL
var i int

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) // use all computer logical CPUs

	go func() {
		for {
			i++
		}
	}()
	go func() {
		for {
			i--
		}
	}()

	for {
	}
}
