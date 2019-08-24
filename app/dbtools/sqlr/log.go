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

func (l logDB) Write(q Query) Result {
	return logRW{rw: l.db, f: l.f}.Write(q)
}

func (l logDB) Read(q Query) Rows {
	return logRW{rw: l.db, f: l.f}.Read(q)
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

func (l logRW) Write(q Query) Result {
	t := time.Now()
	result := l.rw.Write(q)
	l.f(q.String(), result.Rows, t, time.Since(t))
	return result
}

func (l logRW) Read(q Query) Rows {
	t := time.Now()
	rows := l.rw.Read(q)
	l.f(q.String(), -1, t, time.Since(t))
	return rows
}
