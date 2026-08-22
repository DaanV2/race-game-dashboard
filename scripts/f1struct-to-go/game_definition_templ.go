package main

const struct_template = `//WARNING: generate by go script for game_definition.c

package {{.PKG}}

import (
	xbinary "github.com/daanv2/race-game-dashboard/pkg/extensions/binary"
)

{{ $struct_name := .Struct.Name }}
type {{$struct_name}} struct {
	{{range .Struct.Fields}}
	{{.Name}} {{.Type}} // {{.Comment}}
	{{end}}
}

{{range .Struct.Fields}}
func (data *{{$struct_name}}) Get{{.Name}}() {{.Type}} { return data.{{.Name}} }
func (data *{{$struct_name}}) Set{{.Name}}(v {{.Type}}) { data.{{.Name}} = v }
{{end}}

func (data *{{$struct_name}}) Parse(reader xbinary.LittleEndianReader) {
	{{range .Struct.Fields}}
	data.{{.Name}} = reader.Read{{.Type}}()
	{{end}}
}
`
