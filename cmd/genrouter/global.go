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

	retargs := []string{}
	errRetVals := []string{}

	imports := ""

	for _, field := range g.Results.List {
		buf := bytes.NewBuffer([]byte(""))
		format.Node(buf, token.NewFileSet(), field.Type)

		t := string(buf.Bytes())
		switch t {
		case "error":
			errRetVals = append(errRetVals, "errors.New(\"Can't find route\")")
		case "int", "uint", "int32", "uint32", "uint16", "int16", "int8", "uint8", "byte", "char", "uint64", "int64", "float64", "float32", "float":
			errRetVals = append(errRetVals, "0")
		case "string":
			errRetVals = append(errRetVals, "\"\"")
		case "context.Context":
			imports = imports + " \"golang.org/x/net/context\" "
			errRetVals = append(errRetVals, "ctx")
		default:
			errRetVals = append(errRetVals, "nil")
		}

		retargs = append(retargs, t)
	}

	retArgStr := "(" + strings.Join(retargs, ", ") + ")"
	errRetStr := strings.Join(errRetVals, ", ")

	data := globalData{
		Invocation: strings.Join(os.Args[1:], " "),
		Package:    g.pkg.name,
		MapName:    strings.ToLower(fn) + "s",
		LockName:   strings.ToLower(fn) + "s" + "Lock",
		KeyType:    key,
		FnType:     fn,

		AdditionalImports: imports,

		Args:            argStr,
		ReturnParams:    retArgStr,
		ErrorReturnVals: errRetStr,
		CallArgs:        callArgStr,
	}

	t := template.Must(template.New("global").Parse(globalTemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
