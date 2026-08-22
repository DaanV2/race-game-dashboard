package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		panic("expected exactly 2 argument: <pkg> <folder>")
	}

	pkg := args[0]

	dir, err := filepath.Abs(args[1])
	if err != nil {
		panic(err)
	}
	fmt.Println("handling: " + dir)

	creader := &CReader{}

	data, err := os.ReadFile(filepath.Join(dir, "game_definition.c"))
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}
	if string(data) != "" {
		creader.ReadGameDefinition(string(data))
	}

	for _, spec := range creader.Structs {
		exportStruct(dir, pkg, spec)
	}
}

func exportStruct(dir, pkg string, spec *StructSpec) {
	if spec == nil {
		panic("nil spec")
	}
	if spec.Name == "" {
		panic("name of spec is empty")
	}

	filename := structNameToFilename(spec.Name)

	templ, err := template.New("game_struct").Parse(struct_template)
	if err != nil {
		panic(err)
	}

	wr := &strings.Builder{}
	err = templ.Execute(wr, &struct_template_context{
		PKG:    pkg,
		Struct: spec,
	})
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(dir, filename+".auto.go"), []byte(wr.String()), 0)
	if err != nil {
		panic(err)
	}
}

type struct_template_context struct {
	PKG    string
	Struct *StructSpec
}

func structNameToFilename(name string) string {
	return strings.Trim(strings.NewReplacer(
		"DRS", "drs",
		"A", "_a",
		"B", "_b",
		"C", "_c",
		"D", "_d",
		"E", "_e",
		"F", "_f",
		"G", "_g",
		"H", "_h",
		"I", "_i",
		"J", "_j",
		"K", "_k",
		"L", "_l",
		"M", "_m",
		"N", "_n",
		"O", "_o",
		"P", "_p",
		"Q", "_q",
		"R", "_r",
		"S", "_s",
		"T", "_t",
		"U", "_u",
		"V", "_v",
		"W", "_w",
		"X", "_x",
		"Y", "_y",
		"Z", "_z",
	).Replace(name), "_ ")
}
