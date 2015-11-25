package router

//go:generate genrouter -type global -fntype Route -keytype string .

type Payload struct {
	Message string
}

type Route func(p *Payload) error
