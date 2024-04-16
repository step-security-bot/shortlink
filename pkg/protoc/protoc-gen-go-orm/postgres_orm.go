package main

import (
	"path/filepath"
	"strings"

	"google.golang.org/protobuf/compiler/protogen"
)

func generatePostgresFile(gen *protogen.Plugin, file *protogen.File) {
	f := strings.Split(file.GeneratedFilenamePrefix, string(filepath.Separator))
	filename := f[len(f)-1] + ".postgres.orm.go"

	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-orm. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-orm v" + version)
	g.P("// - protoc             ", protocVersion(gen))
	g.P("// source: ", file.Desc.Path())
	g.P()
	g.P("package ", *packageName)
	g.P()

	g.P("import (")
	g.P("\"strings\"")
	g.P("\"github.com/Masterminds/squirrel\"")
	g.P(")")

	processPostgresFile(file, g)
}

func processPostgresFile(file *protogen.File, g *protogen.GeneratedFile) {
	for _, message := range file.Messages {
		// Check if the message should be processed
		if _, ok := filterMap[message.GoIdent.GoName]; ok || len(filterMap) == 0 {
			generatePostgresStructForMessage(message, g)
		}
	}
}

func generatePostgresStructForMessage(message *protogen.Message, g *protogen.GeneratedFile) {
	// Generate the BuildFilter method
	structName := "Filter" + message.GoIdent.GoName

	generateBuildPostgresFilterMethod(g, structName, message.Fields)
}

func generateBuildPostgresFilterMethod(g *protogen.GeneratedFile, structName string, fields []*protogen.Field) {
	g.P("func (f *", structName, ") BuildFilter(query squirrel.SelectBuilder) squirrel.SelectBuilder {")

	for _, field := range fields {
		if field.Desc.IsList() || field.Desc.IsMap() {
			continue
		}
		fieldName := field.GoName
		dbColumnName := strings.ToLower(fieldName)
		generateFieldFilterConditions(g, fieldName, dbColumnName)
	}

	g.P("return query")
	g.P("}")
}

func generateFieldFilterConditions(g *protogen.GeneratedFile, fieldName, dbColumnName string) {
	conditions := []string{"Eq", "Ne", "Lt", "Le", "Gt", "Ge", "StartsWith", "EndsWith"}
	operators := []string{"=", "<>", "<", "<=", ">", ">=", "LIKE", "LIKE"}

	g.P("if f.", fieldName, " != nil {") // nil pointer check

	for i, cond := range conditions {
		op := operators[i]
		valuePlaceholder := "?"
		if cond == "StartsWith" {
			valuePlaceholder = "'%' || ?"
		} else if cond == "EndsWith" {
			valuePlaceholder = "? || '%'"
		}
		g.P("if f.", fieldName, ".", cond, " != \"\" {")
		g.P("query = query.Where(\"", dbColumnName, " ", op, " ", valuePlaceholder, "\", f.", fieldName, ".", cond, ")")
		g.P("}")
	}

	// Handle Contains
	g.P("if len(f.", fieldName, ".Contains) > 0 {")
	g.P("containsQueries := []string{}")
	g.P("containsArgs := []interface{}{}")
	g.P("for _, v := range f.", fieldName, ".Contains {")
	g.P("    if v != \"\" {")
	g.P("        containsQueries = append(containsQueries, \"", dbColumnName, " LIKE ?\")")
	g.P("        containsArgs = append(containsArgs, \"%\" + v + \"%\")")
	g.P("    }")
	g.P("}")
	g.P("if len(containsQueries) > 0 {")
	g.P("    query = query.Where(\"(\" + strings.Join(containsQueries, \" OR \") + \")\", containsArgs...)")
	g.P("}")
	g.P("}")

	// Handle NotContains
	g.P("if len(f.", fieldName, ".NotContains) > 0 {")
	g.P("notContainsQueries := []string{}")
	g.P("notContainsArgs := []interface{}{}")
	g.P("for _, v := range f.", fieldName, ".NotContains {")
	g.P("    if v != \"\" {")
	g.P("        notContainsQueries = append(notContainsQueries, \"", dbColumnName, " NOT LIKE ?\")")
	g.P("        notContainsArgs = append(notContainsArgs, \"%\" + v + \"%\")")
	g.P("    }")
	g.P("}")
	g.P("if len(notContainsQueries) > 0 {")
	g.P("    query = query.Where(\"(\" + strings.Join(notContainsQueries, \" OR \") + \")\", notContainsArgs...)")
	g.P("}")
	g.P("}")

	g.P("if f.", fieldName, ".IsEmpty {")
	g.P("query = query.Where(\"", dbColumnName, " = '' OR ", dbColumnName, " IS NULL\")")
	g.P("}")

	g.P("if f.", fieldName, ".IsNotEmpty {")
	g.P("query = query.Where(\"", dbColumnName, " <> '' AND ", dbColumnName, " IS NOT NULL\")")
	g.P("}")

	g.P("}") // End of nil pointer check
}
