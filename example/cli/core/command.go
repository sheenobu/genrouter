package core

//go:generate genrouter -type context -fntype Command -keytype string .

import (
	"golang.org/x/net/context"
)

type Message struct {
	Name string
	Data []string
}

type Command func(ctx context.Context, msg *Message) error
