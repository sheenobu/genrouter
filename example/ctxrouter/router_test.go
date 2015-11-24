package router

import (
	"golang.org/x/net/context"
	"sync"
	"testing"
)

func TestRouter(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	ctx := RegisterRoute(context.Background(), "my-route", func(ctx context.Context) error {
		wg.Done()
		return nil
	})

	CallRoute(ctx, "my-route")

	wg.Wait()
}
