package helpers

import "regexp"

var (
	carriageReturnRegex       = regexp.MustCompile("\r")
	whitespaceCharactersRegex = regexp.MustCompile("\\s+$")
	onlySpacesStringRegex     = regexp.MustCompile(" +\n")
	headNewlines              = regexp.MustCompile("^\n+")
	tooManyNewlines           = regexp.MustCompile("\n{4,}")
	moderTagsRegex            = regexp.MustCompile("(?i)\\[moder](.*)\\[/moder]")
	quoteTagsRegex            = regexp.MustCompile("(?i)\\[q](.*)\\[/q]")
	nonRussianCharactersRegex = regexp.MustCompile("[^А-Яа-я]")
)

// Чистим сообщение от лишних пробельных символов
func FormatMessage(message string) string {
	result := message

	// удаляем символ возврата каретки
	result = carriageReturnRegex.ReplaceAllLiteralString(result, "")
	// удаляем пробелы и переводы строки в конце сообщения
	result = whitespaceCharactersRegex.ReplaceAllLiteralString(result, "")
	// удаляем пробелы, если строка состоит только из них
	result = onlySpacesStringRegex.ReplaceAllLiteralString(result, "\n")
	// удаляем переводы строки в начале сообщения
	result = headNewlines.ReplaceAllLiteralString(result, "")
	// если переводов строки больше 3, оставляем 3
	result = tooManyNewlines.ReplaceAllLiteralString(result, "\n\n\n")

	return result
}

// Проверяем сообщение на наличие тегов [moder]...[/moder], поиск регистронезависимый
func ContainsModerTags(message string) bool {
	return moderTagsRegex.MatchString(message)
}

// Удаляем теги [moder]...[/moder], поиск регистронезависимый
func RemoveModerTags(message string) string {
	return moderTagsRegex.ReplaceAllString(message, "$1")
}

// Удаляем символы, которые не будут участвовать в вычислении длины сообщения
func RemoveImmeasurable(message string) string {
	result := message

	// удаляем цитаты
	result = quoteTagsRegex.ReplaceAllLiteralString(result, "")
	// удаляем все нерусские символы
	result = nonRussianCharactersRegex.ReplaceAllLiteralString(result, "")

	return result
}
