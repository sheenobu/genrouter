package main

import (
	"os"
	"strings"

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
	data := contextData{
		Invocation:  strings.Join(os.Args[1:], " "),
		Package:     g.pkg.name,
		KeyType:     key,
		FnType:      fn,
		FnTypeLower: strings.ToLower(fn),

		ContextKeyType: strings.ToLower(fn) + "RouterKeyType",
		ContextKey:     strings.ToLower(fn) + "RouterKey",

		//TODO: populate empty values based on parsed type
		Args:            "",
		ReturnParams:    "error",
		ErrorReturnVals: "errors.New(\"Can't find route\")",
		CallArgs:        "",
	}

	t := template.Must(template.New("context").Parse(contextTemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
