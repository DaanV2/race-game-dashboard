package main

type StructSpec struct {
	Name   string
	Fields []*FieldSpec
}

type FieldSpec struct {
	Name    string
	Type    string
	Comment string
}
