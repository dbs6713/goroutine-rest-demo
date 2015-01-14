package main

// This program makes HTTP requests to three different URLs using goroutines.
// We will wait for each goroutine to finish by using a wait group. You can
// think of a goroutine as a fork, and the wait group as a counter that waits
// for n forks to terminate.
//
// What you should see is that slower HTTP calls will appear last in the
// output, even if they appear first in the code. In my case, bbc.co.uk takes
// longest to respond and appears last in the program's output, even though it
// appears first in the code.

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
)

// Get a URL and output the status code response. We will notify the WaitGroup
// that we are done after the status code is output.
func get(url string, w *sync.WaitGroup) {
	defer w.Done()
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(url, res.StatusCode)

}

func main() {
	// This is the number of OS threads Go will use to run our program. We need
	// to set this > 1 to see real concurrency, otherwise our goroutines will
	// just run sequentially in one thread.
	runtime.GOMAXPROCS(4)

	// We need to let go know that it should wait for 3 goroutines to finish
	// before exiting the program. If we don't do this the program will exit
	// before the HTTP calls are made. If we specify too many in Add() our
	// program will hang indefinitely after the HTTP calls are made. Either
	// case is not what we want.
	var w sync.WaitGroup
	w.Add(3)
	go get("http://bbc.co.uk", &w)
	go get("http://google.com", &w)
	go get("http://github.com", &w)
	w.Wait()
}
