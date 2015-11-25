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

	ctx := RegisterRoute(context.Background(), "my-route", func(ctx context.Context, str string) (context.Context, error) {
		if str != "hello world" {
			return ctx, errors.New("Mismatched")
		}
		wg.Done()
		return ctx, nil
	})

	newCtx, err := CallRoute(ctx, "my-route", "hello world")
	if err != nil {
		t.Errorf("Error should have been nil, is %s", err)
	}

	if newCtx == nil {
		t.Errorf("new context should not have been nil")
	}

	wg.Wait()
}
