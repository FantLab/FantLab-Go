package sqlr

type NoRows struct {
	Err error
}

func (rows NoRows) Error() error {
	return rows.Err
}
func (rows NoRows) Scan(output interface{}) error {
	return rows.Err
}
