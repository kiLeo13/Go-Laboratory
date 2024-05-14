package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {

    t.Run("it runs safely concurrently", func(t *testing.T) {

        wantedCount := 1000
        counter := NewCounter()

        var wg sync.WaitGroup
        wg.Add(wantedCount)

        for i := 0; i < wantedCount; i++ {
            go func() {
                counter.Inc()
                wg.Done()
            }()
        }

        wg.Wait()

        assertCounter(t, counter, wantedCount)
    })
}

func NewCounter() *Counter {
    return &Counter{}
}

func assertCounter(t testing.TB, got *Counter, want int) {
    t.Helper()

    if got.Value() != want {
        t.Errorf("got %d, want %d", got.Value(), want)
    }
}

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Inc() {
    c.mu.Lock()
    c.value++
    c.mu.Unlock()
}

func (c *Counter) Value() int {
    return c.value
}