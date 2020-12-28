package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		slowServer := makeDelayServer(20 * time.Second)
		fastServer := makeDelayServer(0 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Error("expected an error but didn't get one")
		}

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		server := makeDelayServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
	// slowServer := makeDelayServer(20 * time.Microsecond)

	// fastServer := makeDelayServer(0 * time.Microsecond)

	// slowUrl := slowServer.URL
	// fastUrl := fastServer.URL

	// want := fastUrl
	// got := Racer(slowUrl, fastUrl)

	// if got != want {
	// 	t.Errorf("got '%s', want '%s'", got, want)
	// }

	// slowServer.Close()
	// fastServer.Close()
}
