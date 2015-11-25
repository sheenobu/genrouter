// generated by genrouter -type global -fntype App -keytype string .; DO NOT EDIT

package core

import (
	"errors"
	"golang.org/x/net/context"
	"sync"
)

var apps map[string]App
var appsLock sync.RWMutex

func init() {
	appsLock.Lock()
	defer appsLock.Unlock()
	apps = make(map[string]App)
}

func RegisterApp(key string, val App) {
	appsLock.Lock()
	defer appsLock.Unlock()
	apps[key] = val
}

func CallApp(key string, ctx context.Context) (context.Context, error) {
	appsLock.RLock()
	defer appsLock.RUnlock()
	r, ok := apps[key]

	if !ok {
		return ctx, errors.New("Can't find route")
	}

	return r(ctx)
}
