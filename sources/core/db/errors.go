package db

import (
	"database/sql"
	"github.com/FantLab/go-kit/database/rowscanner"
)

func IsNotFoundError(err error) bool {
	return err == sql.ErrNoRows || err == rowscanner.ErrInvalidRowCount /* нет результатов при сканировании в скаляр */
}
