package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

    slowServer := makeDelayedServer(20 * time.Millisecond);
    fastServer := makeDelayedServer(0)

    defer slowServer.Close()
    defer fastServer.Close()
    
    slowUrl := slowServer.URL
    fastUrl := fastServer.URL

    want := fastUrl
    got := Racer(slowUrl, fastUrl);

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
    return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        time.Sleep(delay)
        w.WriteHeader(http.StatusOK)
    }))
}

func Racer(a, b string) (winner string) {

    select {
    case <-ping(a):
        return a
    case <-ping(b):
        return b
    }
}

func ping(url string) chan struct{} {
    ch := make(chan struct{})
    go func()  {
        http.Get(url)
        close(ch)
    }()
    return ch;
}