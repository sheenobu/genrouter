package router

import (
	"errors"
	"golang.org/x/net/context"
	"sync"
	"testing"
)

func TestRouter(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	ctx := RegisterRoute(context.Background(), "my-route", func(ctx context.Context, str string) error {
		if str != "hello world" {
			return errors.New("Mismatched")
		}
		wg.Done()
		return nil
	})

	CallRoute(ctx, "my-route", "hello world")

	wg.Wait()
}
