package main

import (
	"bytes"
	"os"
	"strings"

	"go/format"
	"go/token"

	"text/template"
)

type contextData struct {
	Invocation        string
	AdditionalImports string

	Package        string
	KeyType        string
	FnType         string
	FnTypeLower    string
	ContextKeyType string
	ContextKey     string

	Args            string
	ReturnParams    string
	ErrorReturnVals string
	CallArgs        string
}

func (g *Generator) generateContext(key string, fn string) {

	args := []string{}
	callargs := []string{}

	for _, field := range g.Params.List[1:] {

		buf := bytes.NewBuffer([]byte(""))
		format.Node(buf, token.NewFileSet(), field.Type)

		args = append(args, field.Names[0].Name+" "+string(buf.Bytes()))
		callargs = append(callargs, field.Names[0].Name)
	}

	argStr := ""
	if len(args) > 0 {
		argStr = ", " + strings.Join(args, ", ")
	}

	callArgStr := ""
	if len(callargs) > 0 {
		callArgStr = strings.Join(callargs, ", ")
	}

	retargs := []string{}
	errRetVals := []string{}

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
			errRetVals = append(errRetVals, "ctx")
		default:
			errRetVals = append(errRetVals, "nil")
		}

		retargs = append(retargs, t)
	}

	retArgStr := "(" + strings.Join(retargs, ", ") + ")"
	errRetStr := strings.Join(errRetVals, ", ")

	data := contextData{
		Invocation:  strings.Join(os.Args[1:], " "),
		Package:     g.pkg.name,
		KeyType:     key,
		FnType:      fn,
		FnTypeLower: strings.ToLower(fn),

		ContextKeyType: strings.ToLower(fn) + "RouterKeyType",
		ContextKey:     strings.ToLower(fn) + "RouterKey",

		Args:            argStr,
		ReturnParams:    retArgStr,
		ErrorReturnVals: errRetStr,
		CallArgs:        callArgStr,
	}

	t := template.Must(template.New("context").Parse(contextTemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
