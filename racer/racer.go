package racer

import (
	"net/http"
)

func Racer(a, b string) (winner string) {
	// `select` allows you to wait on multiple channels, and the
	// first one wins
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

// an empty struct uses the least amount of memory of any
// datatype, it's used because it's the most lightweight
func ping(url string) chan struct{} {
	// channels must be created with make. create a channel as
	// a var without init will make it nil, which will block
	// forever
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}