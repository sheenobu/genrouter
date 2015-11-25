# genrouter

genrouter is a golang tool that allows the creation of routing systems based on in-memory maps. Maps
can be globally accessable with a RWMutex or they can be attached to an x/net/context.

## Install

	go install github.com/sheenobu/genrouter/cmd/...

## Usage

route.go:

	//go:generate genrouter -type global -fntype MyRoute -keytype string
	
	type MyRoute func() error


main.go:

	func MyCustomRoute() error {
		return nil
	}

	func main() {
		RegisterMyRoute("name", MyCustomRoute)
		CallMyRoute("name")
	}


## License

Code is based on stringer, therefore has the go BSD-style license

## TODO:

 * Generate return values properly using ast
 * Better error names
 * Potentially make an object-based approach as well?
    * r :=NewRouter()
	* r.Register("X", ...)
	* r.Call("X", ...)
 * Configurable Register and Call function names?



