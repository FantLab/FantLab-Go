package sqlr

type Result struct {
	Rows  int64
	Error error
}

type Rows interface {
	Error() error
	Scan(output interface{}) error
}

type ReaderWriter interface {
	Exec(query string, args ...interface{}) Result
	Query(query string, args ...interface{}) Rows
}

type Transactional interface {
	InTransaction(perform func(ReaderWriter) error) error
}

type DB interface {
	Transactional
	ReaderWriter
}
