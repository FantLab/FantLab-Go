package scanr

import "errors"

var (
	ErrNoRows          = errors.New("scanr: no rows")
	ErrMultiRows       = errors.New("scanr: multiple rows found")
	ErrMultiColumns    = errors.New("scanr: multiple columns found")
	ErrUnsupportedType = errors.New("scanr: unsupported output type")
	ErrIsNil           = errors.New("scanr: output value must not be nil")
	ErrNotAPtr         = errors.New("scanr: output value must be a pointer")
	ErrNotAStruct      = errors.New("scanr: slice element type must be a struct")
)
