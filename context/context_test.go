package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
    
    t.Run("returns data from store", func(t *testing.T) {
        data := "Hello, world"
        store := &SpyStore{response: data, t: t}
        svr := Server(store)

        request := httptest.NewRequest(http.MethodGet, "/", nil)

        cancelCtx, cancel := context.WithCancel(request.Context())
        time.AfterFunc(5 * time.Millisecond, cancel)
        request = request.WithContext(cancelCtx)

        response := &SpyResponseWriter{}

        svr.ServeHTTP(response, request)

        if response.written {
            t.Error("a response should not have been written")
        }
    })
}