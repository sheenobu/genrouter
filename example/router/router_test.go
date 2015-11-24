package router

import (
	"sync"
	"testing"
)

func TestRouter(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	RegisterRoute("my-route", func() error {
		wg.Done()
		return nil
	})

	CallRoute("my-route")

	wg.Wait()

}
