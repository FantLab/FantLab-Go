package scanr

import "errors"

var (
	ErrNoRows          = errors.New("No rows")
	ErrMultiRows       = errors.New("Multiple rows found")
	ErrMultiColumns    = errors.New("Multiple columns found")
	ErrUnsupportedType = errors.New("Unsupported output type")
	ErrIsNil           = errors.New("Output value must not be nil")
	ErrNotAPtr         = errors.New("Output value must be a pointer")
	ErrNotAStruct      = errors.New("Slice element type must be a struct")
)
