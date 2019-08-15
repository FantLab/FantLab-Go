package sqlr

import (
	"database/sql"
	"fantlab/scanr"
)

type Rows struct {
	data  *sql.Rows
	Error error
}

func (rows Rows) Scan(output interface{}) error {
	if rows.Error != nil {
		return rows.Error
	}

	return scanr.Scan(output, &sqlRows{
		data:                rows.data,
		nullablesAsDefaults: true,
	})
}
