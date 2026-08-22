package main

import (
	"regexp"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type CReader struct {
	Structs []*StructSpec

	currentSpec  *StructSpec
	currentField *FieldSpec
}

func (c *CReader) startNewStruct() {
	c.finishStruct()

	c.currentSpec = &StructSpec{}
	c.Structs = append(c.Structs, c.currentSpec)
}

func (c *CReader) startNewField() {
	c.finishField()
	if c.currentSpec == nil {
		panic("trying to write new field with no spec")
	}
	c.currentField = &FieldSpec{}
	c.currentSpec.Fields = append(c.currentSpec.Fields, c.currentField)
}

func (c *CReader) finishStruct() {
	c.finishField()

	if c.currentSpec != nil {
		c.currentSpec.Name = strings.TrimSpace(c.currentSpec.Name)
	}

	c.currentSpec = nil
}

func (c *CReader) finishField() {
	if c.currentField != nil {
		c.currentField.Name = strings.TrimSpace(c.currentField.Name)
		c.currentField.Type = strings.TrimSpace(c.currentField.Type)
		c.currentField.Comment = strings.TrimSpace(c.currentField.Comment)
	}
	c.currentField = nil
}

func (c *CReader) ReadGameDefinition(code string) {
	lines := strings.Split(code, "\n")

	for i := range lines {
		// Clean line
		line := strings.TrimSpace(lines[i])

		// Struct defintion start?
		if name, ok := strings.CutPrefix(line, "struct"); ok {
			c.startNewStruct()
			c.currentSpec.Name = TrimBraces(strings.TrimSpace(name))
			continue
		}
		// If we have a spec, check for ending and new fields?
		if c.currentSpec != nil {
			// End of struct?
			if name, ok := strings.CutPrefix(line, "}"); ok {
				if name != ";" {
					c.currentSpec.Name = strings.TrimSpace(strings.Trim(name, ";"))
				}
				c.finishStruct()
			}

			if c.currentField != nil {
				// Field definition?
				reg := regexp.MustCompile("\\w*(.+)\\w*(.+);")
				if reg.MatchString(line) {
					matches := reg.SubexpNames()
					if len(matches) != 2 {
						c.startNewField()
						c.currentField.Name = fixFieldName(matches[1])
						c.currentField.Type = fixFieldType(matches[0])
					}
				}

				// Process comment
				if _, after, ok := strings.Cut(line, "//"); ok {
					c.currentField.Comment += after + " "
				}

			}
		}
	}
}

func fixFieldName(s string) string {
	if after, ok := strings.CutPrefix(s, "m_"); ok {
		s = after
	}

	return cases.Title(language.BritishEnglish).String(s)
}

func fixFieldType(s string) string {
	switch s {
	case "uint8":
		return "uint8" // Unsigned 8-bit integer
	case "int8":
		return "int8" // Signed 8-bit integer
	case "uint16":
		return "uint16" // Unsigned 16-bit integer
	case "int16":
		return "int16" // Signed 16-bit integer
	case "uint32":
		return "uint32" // Unsigned 32-bit integer
	case "float":
		return "float32" // Floating point (32-bit)
	case "double":
		return "float64" // Double-precision floating point (64-bit)
	case "uint64":
		return "uint64" // Unsigned 64-bit integer
	case "char":
		return "byte" // Character
	}

	return s
}

func TrimBraces(s string) string {
	return strings.Trim(s, "{}")
}
