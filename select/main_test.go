package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("return faster url", func(t *testing.T) {
		slowServer := makeDelayServer(30 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := ConfigurableRacer(slowURL, fastURL, 10*time.Second)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if want != got {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server does not respond within 10s", func(t *testing.T) {
		serverA := makeDelayServer(2 * time.Second)
		defer serverA.Close()

		serverB := makeDelayServer(3 * time.Second)
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 1*time.Second)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
	}))

	return server
}
