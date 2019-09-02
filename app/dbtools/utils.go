package dbtools

import (
	"database/sql"
	"fantlab/dbtools/scanr"
)

func IsNotFoundError(err error) bool {
	return err == sql.ErrNoRows || err == scanr.ErrNoRows
}
