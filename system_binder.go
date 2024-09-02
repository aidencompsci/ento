package ento

import (
	"reflect"
	"strings"
)

type systemBinder struct {
	System
	queries []queryBinder
	tag     string
}

var queryType reflect.Type

func newSystemBinder(world *World, system System, tag string) systemBinder {
	if queryType == nil {
		queryType = reflect.TypeOf(Query[any]{})
	}
	queries := []queryBinder{}

	systemValues := reflect.ValueOf(system).Elem()
	systemType := systemValues.Type()
	systemFieldsNum := systemValues.NumField()

	for i := 0; i < systemFieldsNum; i++ {
		field := systemType.Field(i)
		if isQuery(field.Type) {
			queryFieldValue := systemValues.Field(i)
			queryValues := queryFieldValue.FieldByIndex([]int{0})
			queryType := queryValues.Type()

			qbind := newQueryBinder(world, queryValues, queryType)
			queries = append(queries, qbind)

			worldBindValue := queryFieldValue.FieldByIndex([]int{1})
			worldBindValue.Set(reflect.ValueOf(world))

			queryBindValue := queryFieldValue.FieldByIndex([]int{2})
			queryBindValue.Set(reflect.ValueOf(&qbind))
		}
	}

	return systemBinder{
		System:  system,
		tag:     tag,
		queries: queries,
	}
}

func isQuery(a reflect.Type) bool {
	if strings.HasPrefix(a.String(), "ento.Query[") {
		return true
	}
	return false
}
