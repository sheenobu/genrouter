# genrouter

genrouter is a golang tool that allows the creation of routing systems based on in-memory maps. Maps
can be globally accessable with a RWMutex or they can be attached to an x/net/context.

## Install

	go install github.com/sheenobu/genrouter/cmd/...

## Usage

route.go:

	import (
		"golang.org/x/net/context"
	)

	//go:generate genrouter -type global -fntype MyRoute -keytype string
	//go:generate genrouter -type context -fntype OtherRoute -keytype string
	
	type MyRoute func(param string) error

	// context must be first!
	type OtherRoute func(ctx context.Context, msg *Message) error

	type Message struct {
		M string
	}

main.go:

	import (
		"golang.org/x/net/context"
	)

	func MyCustomRoute(param string) error {
		return nil
	}

	func MyOtherRoute(ctx context.Context, msg *Message) error {
		return nil
	}

	func main() {
		// routing data is global
		RegisterMyRoute("name", MyCustomRoute)
		CallMyRoute("name", "Hello World")


		// routing data is attached to x/net/context
		ctx := context.Background()
		ctx = RegisterOtherRoute(ctx, "name", MyOtherRoute)

		CallOtherRoute(ctx, "name", &Message{"Hello World"})
	}


## License

Code is based on stringer, therefore has the go BSD-style license

## TODO:

 * Import detection
 * Better error names
 * Potentially make an object-based approach as well?
    * r :=NewRouter()
	* r.Register("X", ...)
	* r.Call("X", ...)
 * Configurable Register and Call function names?



