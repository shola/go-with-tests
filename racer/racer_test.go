package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server{
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0)
	
		// always cleanup resources as close to their definition as possible
		defer slowServer.Close()
		defer fastServer.Close()
	
		slowURL := slowServer.URL
		fastURL := fastServer.URL
	
		want := fastURL
		got, err := Racer(slowURL, fastURL)
	
		if err != nil {
			t.Errorf("got error %q, expected nil", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond before timeout", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)
		timeout := 20 * time.Millisecond

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, timeout)

		if err == nil {
			t.Error("expected an error but didn't get one")
		} 
	})
}