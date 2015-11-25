package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"net/http"
	"net/http/httputil"

	"golang.org/x/net/context"

	"github.com/sheenobu/genrouter/example/cli/core"
)

func lsCommand(ctx context.Context, msg *core.Message) error {
	log.Printf("hello.txt")
	log.Printf("goodbye.txt")
	return nil
}

func echoCommand(ctx context.Context, msg *core.Message) error {

	for _, d := range msg.Data {
		log.Printf("%v", d)
	}

	return nil
}

func getHttp(ctx context.Context, msg *core.Message) error {
	url := strings.TrimSpace(msg.Data[0])
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", dump)
	return nil
}

func headHttp(ctx context.Context, msg *core.Message) error {
	url := strings.TrimSpace(msg.Data[0])
	resp, err := http.Head(url)
	if err != nil {
		return err
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", dump)
	return nil
}

func ShellApp(ctx context.Context) (context.Context, error) {
	ctx = core.RegisterCommand(ctx, "ls", lsCommand)
	ctx = core.RegisterCommand(ctx, "echo", echoCommand)

	return ctx, nil
}

func HttpApp(ctx context.Context) (context.Context, error) {
	ctx = core.RegisterCommand(ctx, "get", getHttp)
	ctx = core.RegisterCommand(ctx, "head", headHttp)
	return ctx, nil
}

func main() {

	running := true

	quit := func(ctx context.Context, msg *core.Message) error {
		log.Printf("quitting")
		running = false
		return nil
	}

	core.RegisterApp("shell", ShellApp)
	core.RegisterApp("http", HttpApp)

	var app string
	var ctx context.Context
	var setup func(string) context.Context

	setApp := func(c2 context.Context, msg *core.Message) error {
		ctx = setup(strings.TrimSpace(msg.Data[0]))
		return nil
	}

	setup = func(appName string) context.Context {
		ctx := context.Background()
		ctx = core.RegisterCommand(ctx, "quit", quit)
		ctx = core.RegisterCommand(ctx, "app", setApp)
		ctx, err := core.CallApp(appName, ctx)
		if err != nil {
			log.Printf("error calling app: %s", err)
		}
		app = appName
		return ctx
	}

	ctx = setup("shell")

	r := bufio.NewReader(os.Stdin)
	w := os.Stdout

	for running {
		w.Write([]byte(fmt.Sprintf("%s > ", app)))
		line, err := r.ReadString('\n')
		if err != nil {
			log.Fatalf("Error: %s", err)
			break
		}
		segments := strings.Split(line, " ")

		cmdName := strings.TrimSpace(segments[0])
		if cmdName == "" {
			continue
		}

		err = core.CallCommand(ctx, cmdName, &core.Message{
			Data: segments[1:],
		})
		if err != nil {
			log.Printf("Error calling command %s: %s", cmdName, err)
		}
	}
}
