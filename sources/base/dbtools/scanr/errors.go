package scanr

import "errors"

var (
	ErrNoRows             = errors.New("scanr: no rows")
	ErrInvalidRowCount    = errors.New("scanr: invalid row count")
	ErrInvalidColumnCount = errors.New("scanr: invalid column count")
	ErrUnsupportedType    = errors.New("scanr: unsupported type")
	ErrIsNil              = errors.New("scanr: output value must not be nil")
	ErrNotAPtr            = errors.New("scanr: output value must be a pointer")
)
