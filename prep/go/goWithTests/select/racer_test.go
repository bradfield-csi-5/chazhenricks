package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("test fast server is returned", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		want := fastUrl
		got, _ := Racer(slowUrl, fastUrl)

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

	})

	t.Run("error if longer than 10 sec", func(t *testing.T) {
		serverA := makeDelayedServer(25 * time.Millisecond)
    timeout := 20 * time.Millisecond

		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL,timeout)
		if err == nil {
			t.Error("expected an error but didnt get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {

	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
