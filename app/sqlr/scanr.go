package sqlr

import (
	"database/sql"
	"fantlab/scanr"
	"reflect"
	"strings"
)

// *******************************************************

type sqlColumn struct {
	name               string
	useNonNullSubField bool
}

func (column *sqlColumn) Name() string {
	return column.name
}

func (column *sqlColumn) Get(value reflect.Value) reflect.Value {
	if column.useNonNullSubField {
		return value.Elem().Field(0)
	}

	return value.Elem()
}

// *******************************************************

type sqlRows struct {
	data                *sql.Rows
	nullablesAsDefaults bool
}

func (rows sqlRows) AltNameTag() string {
	return "db"
}

func (rows sqlRows) IterateUsing(fn scanr.RowFunc) error {
	defer rows.data.Close()

	columnTypes, err := rows.data.ColumnTypes()

	if err != nil {
		return err
	}

	values, columns := getColumnData(columnTypes, rows.nullablesAsDefaults)

	for rows.data.Next() {
		err = rows.data.Scan(values...)

		if err != nil {
			return err
		}

		err = fn(columns, values)

		if err != nil {
			return err
		}
	}

	return rows.data.Err()
}

func getColumnData(columnTypes []*sql.ColumnType, nullablesAsDefaults bool) ([]interface{}, []scanr.Column) {
	size := len(columnTypes)

	values := make([]interface{}, size)
	columns := make([]scanr.Column, size)

	for i, columnType := range columnTypes {
		values[i] = reflect.New(columnType.ScanType()).Interface()

		isNullable := strings.HasPrefix(columnType.ScanType().Name(), "Null")

		columns[i] = &sqlColumn{
			name:               columnType.Name(),
			useNonNullSubField: isNullable && nullablesAsDefaults,
		}
	}

	return values, columns
}
