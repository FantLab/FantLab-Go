package sqlr

import "database/sql"

type Rows struct {
	data  *sql.Rows
	Error error
}

func (rows Rows) Scan(output interface{}) error {
	if rows.Error != nil {
		return rows.Error
	}

	return scanRows(output, rows.data, true)
}
