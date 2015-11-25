package router

import (
	"golang.org/x/net/context"
)

//go:generate genrouter -type context -fntype Route -keytype string .

type Route func(ctx context.Context, name string) (context.Context, error)
