package main

import (
	"os"
	"strings"

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
	data := globalData{
		Invocation: strings.Join(os.Args[1:], " "),
		Package:    g.pkg.name,
		MapName:    strings.ToLower(fn) + "s",
		LockName:   strings.ToLower(fn) + "s" + "Lock",
		KeyType:    key,
		FnType:     fn,

		//TODO: populate empty values based on parsed type
		Args:            "",
		ReturnParams:    "error",
		ErrorReturnVals: "errors.New(\"Can't find route\")",
		CallArgs:        "",
	}

	t := template.Must(template.New("global").Parse(globalTemplate))

	err := t.Execute(&g.buf, data)
	if err != nil {
		panic(err)
	}
}
