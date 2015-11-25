package main

import (
	"bytes"
	"os"
	"strings"

	"go/format"
	"go/token"

	"text/template"
)

type globalData struct {
	Invocation        string
	AdditionalImports string

	Package  string
	MapName  string
	LockName string
	KeyType  string
	FnType   string

	Args            string
	ReturnParams    string
	ErrorReturnVals string
	CallArgs        string
}

func (g *Generator) generateGlobal(key string, fn string) {
	args := []string{}
	callargs := []string{}

	for _, field := range g.Params.List {

		buf := bytes.NewBuffer([]byte(""))
		format.Node(buf, token.NewFileSet(), field.Type)

		args = append(args, field.Names[0].Name+" "+string(buf.Bytes()))
		callargs = append(callargs, field.Names[0].Name)
	}

	argStr := ""
	if len(args) > 0 {
		argStr = strings.Join(args, ", ")
	}

	callArgStr := ""
	if len(callargs) > 0 {
		callArgStr = strings.Join(callargs, ", ")
	}

	data := globalData{
		Invocation: strings.Join(os.Args[1:], " "),
		Package:    g.pkg.name,
		MapName:    strings.ToLower(fn) + "s",
		LockName:   strings.ToLower(fn) + "s" + "Lock",
		KeyType:    key,
		FnType:     fn,

		Args:            argStr,
		ReturnParams:    "error",
		ErrorReturnVals: "errors.New(\"Can't find route\")",
		CallArgs:        callArgStr,
	}

	t := template.Must(template.New("global").Parse(globalTemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
