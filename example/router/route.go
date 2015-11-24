package router

//go:generate genrouter -type global -fntype Route -keytype string .

type Route func() error
