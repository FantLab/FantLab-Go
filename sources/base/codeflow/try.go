package codeflow

func Try(fs ...func() error) error {
	for _, f := range fs {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}
