// generated by genrouter -type context -fntype Command -keytype string .; DO NOT EDIT

package core

import (
	"errors"

	"golang.org/x/net/context"
)

func RegisterCommand(ctx context.Context, key string, val Command) context.Context {
	mp, ok := ctx.Value(commandRouterKey).(map[string]Command)
	if !ok {
		mp = make(map[string]Command)
	}
	mp[key] = val
	return context.WithValue(ctx, commandRouterKey, mp)
}

func CallCommand(ctx context.Context, key string, msg *Message) error {
	r, ok := commandfromContext(ctx, key)
	if !ok {
		return errors.New("Can't find route")
	}

	return r(ctx, msg)
}

type commandRouterKeyType int

var commandRouterKey commandRouterKeyType

func commandfromContext(ctx context.Context, key string) (Command, bool) {
	mp, ok := ctx.Value(commandRouterKey).(map[string]Command)
	if !ok {
		return nil, false
	}

	r, ok := mp[key]

	if !ok {
		return nil, false
	}

	return r, true
}
