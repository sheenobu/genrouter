package router

import (
	"sync"
	"testing"
)

func TestRouter(t *testing.T) {

	var wg sync.WaitGroup

	wg.Add(1)

	RegisterRoute("my-route", func(pl *Payload) error {
		wg.Done()
		return nil
	})

	CallRoute("my-route", &Payload{
		Message: "hello world",
	})

	wg.Wait()

}
