package main

// This program makes HTTP requests to three different URLs using goroutines.
// We will wait for each goroutine to finish by using channels. Each response
// will be sent to the channel, and we continue to read responses from the
// channel until we have received as many as we requested (3 in this case).
//
// What you should see is that slower HTTP calls will appear last in the
// output, even if they appear first in the code. In my case, bbc.co.uk takes
// longest to respond and appears last in the program's output, even though it
// appears first in the code.

import (
	"fmt"
	"net/http"
	"runtime"
)

// Here we create a struct so we can send the response back along a channel. We
// could simply print the response inside get() but it's more illustrative to
// print the response by reading it from the channel.
type response struct {
	url  string
	code int
}

func get(url string, ch chan response) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	ch <- response{url, res.StatusCode}
}

func main() {
	// This is the number of OS threads Go will use to run our program. We need
	// to set this > 1 to see real concurrency, otherwise our goroutines will
	// just run sequentially in one thread.
	runtime.GOMAXPROCS(4)

	var urls = []string{
		"http://bbc.co.uk",
		"http://google.com",
		"http://github.com",
	}

	// We will send responses to this channel.
	ch := make(chan response)

	// Queue up the requests
	for _, url := range urls {
		go get(url, ch)
	}

	// We started three goroutines and asked them to send their output to ch.
	// <-ch will block and wait for one item to be ready, so we can display it.
	// Since <-ch only waits for one item, we will need to use a loop to wait
	// for each item in urls or the program will terminate too early. We could
	// loop for len(urls) times but it's simpler and more idiomatic to use
	// `range urls` to accomplish this.
	for range urls {
		res := <-ch
		fmt.Println(res.url, res.code)
	}
}
