package automapper

import (
	"fmt"
	"reflect"
	"strings"
)

func Generate(automappers []AutoMapper, conf *AutoMapperConf) {
	for _, mapper := range automappers {
		modelType := reflect.TypeOf(mapper.Model.Type)
		transferObjectType := reflect.TypeOf(mapper.Schema.Type)

		fmt.Printf("%s: %s -> %s\n", mapper.Name, modelType.Name(), transferObjectType.Name())

		// From Fields
		mapper.Imports = append(mapper.Imports, modelType.PkgPath())
		mapper.Model.Reference = modelType.Name()
		mapper.Model.Fields = make([]reflect.StructField, 0)
		for i := 0; i < modelType.NumField(); i++ {
			mapper.Model.Fields = append(mapper.Model.Fields, modelType.Field(i))
		}

		// To Fields
		mapper.Imports = append(mapper.Imports, transferObjectType.PkgPath())
		mapper.Schema.Reference = transferObjectType.Name()
		mapper.Schema.Fields = make([]reflect.StructField, 0)
		for i := 0; i < transferObjectType.NumField(); i++ {
			mapper.Schema.Fields = append(mapper.Schema.Fields, transferObjectType.Field(i))
		}

		// Determine Field Assignments by matching the To fields and From fields by name
		mapper.FieldAssignments = make([]FieldAssignment, 0)

		for _, toField := range mapper.Schema.Fields {
			for _, fromField := range mapper.Model.Fields {
				if strings.EqualFold(toField.Name, fromField.Name) {
					mapper.FieldAssignments = append(mapper.FieldAssignments, FieldAssignment{
						ModelField: fromField.Name,
						SchemaField:   toField.Name,
					})
				}
			}
		}

		mapper.ExecuteTemplates(conf)
	}
}
