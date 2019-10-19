package dbtools

import (
	"fantlab/dbtools/scanr"
)

func IsNotFoundError(err error) bool {
	return err == scanr.ErrNoRows
}
