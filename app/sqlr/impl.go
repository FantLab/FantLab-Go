package sqlr

import "time"

type impl struct {
	rw    dbReaderWriter
	logFn LogFunc
}

func (i impl) Exec(query string, args ...interface{}) Result {
	t := time.Now()

	result, err := i.rw.Exec(query, args...)
	rowsAffected, _ := result.RowsAffected()

	i.logFn(formatQuery(query, bindVarChar, args...), rowsAffected, t, time.Since(t))

	return Result{
		RowsAffected: rowsAffected,
		Error:        err,
	}
}

func (i impl) Query(query string, args ...interface{}) Rows {
	t := time.Now()

	rows, err := i.rw.Query(query, args...)

	i.logFn(formatQuery(query, bindVarChar, args...), -1, t, time.Since(t))

	return Rows{
		data:  rows,
		Error: err,
	}
}

func (i impl) QueryIn(query string, args ...interface{}) Rows {
	newQuery, newArgs, err := rebindQuery(query, bindVarChar, args...)

	if err != nil {
		return Rows{
			data:  nil,
			Error: err,
		}
	}

	return i.Query(newQuery, newArgs...)
}
