package sqlr

type Result struct {
	RowsAffected int64
	Error        error
}
