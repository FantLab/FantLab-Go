package sqlr

type Result struct {
	Rows  int64
	Error error
}

type Rows interface {
	Error() error
	Scan(output interface{}) error
}

type Reader interface {
	Read(q Query) Rows
}

type Writer interface {
	Write(q Query) Result
}

type ReaderWriter interface {
	Reader
	Writer
}

type Transactional interface {
	InTransaction(perform func(ReaderWriter) error) error
}

type DB interface {
	Transactional
	ReaderWriter
}
