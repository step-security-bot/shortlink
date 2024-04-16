package main

import (
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func generateCommonFile(gen *protogen.Plugin, file *protogen.File) {
	name := strings.Split(file.GeneratedFilenamePrefix, "/")
	name = append(name[:len(name)-1], commonFilename)
	filename := strings.Join(name, "/")
	g := gen.NewGeneratedFile(filename, "")
	g.P("// Code generated by protoc-gen-go-orm. DO NOT EDIT.")
	g.P("package ", file.GoPackageName) // Use the package name of the first file
	g.P()

	// Generate the StringFilterInput struct definition
	g.P("type StringFilterInput struct { //nolint:unused")
	g.P("    Eq           string    `json:\"eq,omitempty\"`")          // Equal
	g.P("    Ne           string    `json:\"ne,omitempty\"`")          // Not Equal
	g.P("    Lt           string    `json:\"lt,omitempty\"`")          // Less Than
	g.P("    Le           string    `json:\"le,omitempty\"`")          // Less Than or Equal
	g.P("    Gt           string    `json:\"gt,omitempty\"`")          // Greater Than
	g.P("    Ge           string    `json:\"ge,omitempty\"`")          // Greater Than or Equal
	g.P("    Contains     []string  `json:\"contains,omitempty\"`")    // Contains
	g.P("    NotContains  []string  `json:\"notContains,omitempty\"`") // Not Contains
	g.P("    StartsWith   string    `json:\"startsWith,omitempty\"`")  // Starts With
	g.P("    EndsWith     string    `json:\"endsWith,omitempty\"`")    // Ends With
	g.P("    IsEmpty      bool      `json:\"isEmpty,omitempty\"`")     // Is Empty
	g.P("    IsNotEmpty   bool      `json:\"isNotEmpty,omitempty\"`")  // Is Not Empty
	g.P("}")
	g.P()

	for _, message := range file.Messages {
		// Generate the BuildFilter method
		structName := "Filter" + message.GoIdent.GoName
		g.P("type ", structName, " struct {")
		for _, field := range message.Fields {
			if field.Desc.IsList() || field.Desc.IsMap() {
				continue
			}

			fieldName := field.GoName
			goType := "*" + "StringFilterInput"
			g.P(fieldName, " ", goType, " `json:\"", strings.ToLower(fieldName), "\"`")
		}
		g.P("}")
		g.P()
	}
}
