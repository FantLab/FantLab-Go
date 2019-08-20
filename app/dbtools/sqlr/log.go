package sqlr

import "time"

type LogFunc func(query string, rows int64, time time.Time, duration time.Duration)

func Log(db DB, f LogFunc) DB {
	return &logDB{db: db, f: f}
}

// *******************************************************

type logDB struct {
	db DB
	f  LogFunc
}

func (l logDB) Exec(query string, args ...interface{}) Result {
	return logRW{rw: l.db, f: l.f}.Exec(query, args...)
}

func (l logDB) Query(query string, args ...interface{}) Rows {
	return logRW{rw: l.db, f: l.f}.Query(query, args...)
}

func (l logDB) InTransaction(perform func(ReaderWriter) error) error {
	return l.db.InTransaction(func(rw ReaderWriter) error {
		return perform(logRW{rw: rw, f: l.f})
	})
}

// *******************************************************

type logRW struct {
	rw ReaderWriter
	f  LogFunc
}

func (l logRW) Exec(query string, args ...interface{}) Result {
	t := time.Now()
	result := l.rw.Exec(query, args...)
	l.f(FormatQuery(query, args...), result.Rows, t, time.Since(t))
	return result
}

func (l logRW) Query(query string, args ...interface{}) Rows {
	t := time.Now()
	rows := l.rw.Query(query, args...)
	l.f(FormatQuery(query, args...), -1, t, time.Since(t))
	return rows
}
