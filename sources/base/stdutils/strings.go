package stdutils

func Elvis(value string, fallback string) string {
	if "" == value {
		return fallback
	}
	return value
}
