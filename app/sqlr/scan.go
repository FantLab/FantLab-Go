package sqlr

import (
	"database/sql"
	"errors"
	"reflect"
	"strings"
)

type columnInfo struct {
	name        string
	hasNullType bool
}

type rowMapFunc func(columns []*columnInfo, values []interface{}) error

func mapRowsUsingFunc(rows *sql.Rows, mapFn rowMapFunc) error {
	defer rows.Close()

	columnTypes, err := rows.ColumnTypes()

	if err != nil {
		return err
	}

	columnsCount := len(columnTypes)

	values := make([]interface{}, columnsCount)
	columns := make([]*columnInfo, columnsCount)

	for i, columnType := range columnTypes {
		values[i] = reflect.New(columnType.ScanType()).Interface()
		columns[i] = &columnInfo{
			name:        columnType.Name(),
			hasNullType: strings.HasPrefix(columnType.ScanType().Name(), "Null"),
		}
	}

	for rows.Next() {
		err = rows.Scan(values...)

		if err != nil {
			return err
		}

		err = mapFn(columns, values)

		if err != nil {
			return err
		}
	}

	return rows.Err()
}

func setSingleValue(value interface{}, output reflect.Value, takeNonNullSubField bool) {
	x := reflect.ValueOf(value).Elem()

	if takeNonNullSubField {
		x = x.Field(0)
	}

	output.Set(x.Convert(output.Type()))
}

func setValuesToStruct(values []interface{}, columns []*columnInfo, output reflect.Value, idxMap map[string]int, nullablesAsDefaults bool) {
	for i, value := range values {
		column := columns[i]

		j, ok := idxMap[column.name]

		if !ok {
			continue
		}

		setSingleValue(value, output.Field(j), nullablesAsDefaults && column.hasNullType)
	}
}

func makeFieldNameIndexMapFromStruct(t reflect.Type) map[string]int {
	m := make(map[string]int)

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		m[f.Name] = i

		if dbName, ok := f.Tag.Lookup("db"); ok {
			m[dbName] = i
		}
	}

	return m
}

func scanRowsIntoSliceOfStructs(output reflect.Value, elemType reflect.Type, rows *sql.Rows, nullablesAsDefaults bool) error {
	idxMap := makeFieldNameIndexMapFromStruct(elemType)

	err := mapRowsUsingFunc(rows, func(columns []*columnInfo, values []interface{}) error {
		newElem := reflect.Indirect(reflect.New(elemType))

		setValuesToStruct(values, columns, newElem, idxMap, nullablesAsDefaults)

		output.Set(reflect.Append(output, newElem))

		return nil
	})

	return err
}

func scanSingleRowIntoStruct(output reflect.Value, rows *sql.Rows, nullablesAsDefaults bool) error {
	idxMap := makeFieldNameIndexMapFromStruct(output.Type())

	once := false

	err := mapRowsUsingFunc(rows, func(columns []*columnInfo, values []interface{}) error {
		if once {
			return errors.New("Multiple rows were found")
		}

		setValuesToStruct(values, columns, output, idxMap, nullablesAsDefaults)

		once = true

		return nil
	})

	if !once {
		return sql.ErrNoRows
	}

	return err
}

func scanSingleValue(output reflect.Value, rows *sql.Rows, nullablesAsDefaults bool) error {
	once := false

	err := mapRowsUsingFunc(rows, func(columns []*columnInfo, values []interface{}) error {
		if len(values) != 1 {
			return errors.New("Multiple columns found")
		}

		if once {
			return errors.New("Multiple rows found")
		}

		setSingleValue(values[0], output, nullablesAsDefaults && columns[0].hasNullType)

		once = true

		return nil
	})

	if !once {
		return sql.ErrNoRows
	}

	return err
}

func scanRowsIntoValue(output reflect.Value, rows *sql.Rows, nullablesAsDefaults bool) error {
	switch output.Type().Kind() {
	case reflect.Slice:
		elemType := output.Type().Elem()

		if elemType.Kind() != reflect.Struct {
			return errors.New("Slice element type must be a struct")
		}

		return scanRowsIntoSliceOfStructs(output, elemType, rows, nullablesAsDefaults)
	case reflect.Struct:
		return scanSingleRowIntoStruct(output, rows, nullablesAsDefaults)
	case
		reflect.Bool,
		reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Float32,
		reflect.Float64,
		reflect.String:
		return scanSingleValue(output, rows, nullablesAsDefaults)
	default:
		return errors.New("Unsupported output type")
	}
}

func scanRows(output interface{}, rows *sql.Rows, nullablesAsDefaults bool) error {
	value := reflect.ValueOf(output)

	if value.Kind() != reflect.Ptr {
		return errors.New("Output value must be a pointer")
	}

	if value.IsNil() {
		return errors.New("Output value must not be nil")
	}

	value = reflect.Indirect(value)

	err := scanRowsIntoValue(value, rows, nullablesAsDefaults)

	if err != nil {
		value.Set(reflect.Zero(value.Type()))
	}

	return err
}
