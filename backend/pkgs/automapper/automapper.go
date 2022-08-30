package automapper

import (
	"bytes"
	"fmt"
	"go/format"
	"os"
	"reflect"
	"strings"
	"text/template"
)

type FieldAssignment struct {
	ModelField  string
	SchemaField string
}

type Model struct {
	Type      interface{}
	Prefix    string
	Fields    []reflect.StructField
	Reference string
}

type Schema struct {
	Name      string
	Type      interface{}
	Prefix    string
	Fields    []reflect.StructField
	Reference string
}

type AutoMapper struct {
	Name             string
	Package          string
	Prefix           string
	Schema           Schema
	Model            Model
	Imports          []string
	FieldAssignments []FieldAssignment
}

func (mapper *AutoMapper) ExecuteTemplates(conf *AutoMapperConf) {
	t := template.New("automapper")
	t, err := t.Parse(automapperTemplate)
	if err != nil {
		fmt.Println(err)
	}

	// Ensure the output directory exists
	os.MkdirAll(conf.OutDir, 0755)

	var path = fmt.Sprintf("%s/%s", conf.OutDir, mapper.GetFileName())

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	var buf bytes.Buffer

	err = t.Execute(&buf, mapper)
	if err != nil {
		fmt.Println(err)
	}

	text, err := format.Source(buf.Bytes())

	if err != nil {
		fmt.Println(err)
	}

	f.Write(text)

}

// GetFileName returns the computed file name based off user preference.
// If the Prefix has been specified on the AutoMapper it will be used
// in place of the Struct name. If the Prefix is not specified, the
// Struct name will be used.
//
// Examples:
// prefix_automapper.go
// mystructname_automapper.go
func (mapper *AutoMapper) GetFileName() string {
	if mapper.Prefix == "" {
		return strings.ToLower(mapper.Schema.Reference) + "_" + "automapper.go"
	}
	return strings.ToLower(mapper.Prefix) + "_" + "automapper.go"

}
