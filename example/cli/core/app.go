package core

//go:generate genrouter -type global -fntype App -keytype string .

import (
	"golang.org/x/net/context"
)

type App func(ctx context.Context) (context.Context, error)
