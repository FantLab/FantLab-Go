package sqlbuilder

import (
	"errors"
	"fantlab/base/dbtools/sqlr"
	"reflect"
	"strings"
)

var (
	ErrInsertNoData             = errors.New("sqlbuilder: insert no data")
	ErrInsertNonHomogeneousData = errors.New("sqlbuilder: insert non homogeneous data")
	ErrInsertUnsupportedType    = errors.New("sqlbuilder: insert unsupported type")
)

func InsertInto(tableName string, entries ...interface{}) sqlr.Query {
	query, err := insertInto(tableName, "db", entries...)
	if err != nil {
		panic(err)
	}
	return *query
}

func insertInto(tableName, tagName string, entries ...interface{}) (*sqlr.Query, error) {
	var typ reflect.Type
	for _, entry := range entries {
		if typ == nil {
			typ = reflect.TypeOf(entry)
		} else if typ != reflect.TypeOf(entry) {
			return nil, ErrInsertNonHomogeneousData
		}
	}
	if typ == nil {
		return nil, ErrInsertNoData
	}
	if typ.Kind() != reflect.Struct {
		return nil, ErrInsertUnsupportedType
	}

	var fieldNames []string

	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		name := f.Tag.Get(tagName)
		if "" == name {
			name = f.Name
		}

		fieldNames = append(fieldNames, name)
	}

	if fieldNames == nil {
		return nil, ErrInsertUnsupportedType
	}

	var sb strings.Builder

	sb.WriteString("INSERT INTO ")
	sb.WriteString(tableName)
	sb.WriteRune('(')
	for i, fieldName := range fieldNames {
		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteString(fieldName)
	}
	sb.WriteRune(')')
	sb.WriteString(" VALUES ")

	values := make([]interface{}, 0, len(entries)*len(fieldNames))

	for i, entry := range entries {
		value := reflect.ValueOf(entry)

		if i > 0 {
			sb.WriteRune(',')
		}
		sb.WriteRune('(')
		for j := 0; j < len(fieldNames); j++ {
			values = append(values, value.Field(j).Interface())

			if j > 0 {
				sb.WriteRune(',')
			}
			sb.WriteRune('?')
		}
		sb.WriteRune(')')
	}

	query := sqlr.NewQuery(sb.String()).WithArgs(values...)

	return &query, nil
}
