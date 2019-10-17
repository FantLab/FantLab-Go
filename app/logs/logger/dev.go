package logger

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	colorRed    = 31
	colorGreen  = 32
	colorYellow = 33
	// colorBlue    = 34
	colorMagenta = 35
	colorCyan    = 36
	colorWhite   = 37

	colorBold     = 1
	colorDarkGray = 90
)

func colorize(s interface{}, c int) string {
	return fmt.Sprintf("\x1b[%dm%v\x1b[0m", c, s)
}

func formatDuration(d time.Duration) string {
	return fmt.Sprintf("%.2fms", float64(d.Nanoseconds()/1e4)/100.0)
}

func statusCodeColor(statusCode int) int {
	switch statusCode / 100 {
	case 4, 5:
		return colorRed
	case 3:
		return colorYellow
	case 2:
		return colorGreen
	case 1:
		return colorDarkGray
	default:
		return colorWhite
	}
}

func DebugString(httpData HTTPData, entries []Entry) string {
	var sb strings.Builder

	sb.WriteString(colorize(httpData.Time.Format("15:04:05"), colorDarkGray))
	sb.WriteRune(' ')

	sb.WriteString(colorize(httpData.Request.Method, colorCyan))
	sb.WriteRune(' ')

	sb.WriteString(colorize(httpData.Request.RequestURI, colorBold))
	sb.WriteRune(' ')

	if httpData.StatusCode > 0 {
		sb.WriteString(colorize(strconv.Itoa(httpData.StatusCode), statusCodeColor(httpData.StatusCode)))
		sb.WriteRune(' ')
	}

	sb.WriteString(colorize(formatDuration(httpData.Duration), colorMagenta))
	sb.WriteRune('\n')

	for i, entry := range entries {
		sb.WriteString(colorize(strconv.Itoa(i+1)+") ", colorYellow))

		if "" != entry.Message {
			sb.WriteString(colorize(entry.Message, colorBold))
			sb.WriteRune(' ')
		}

		for key, value := range entry.More {
			sb.WriteString(colorize(key, colorDarkGray))
			sb.WriteRune(' ')
			sb.WriteString(colorize(value, colorWhite))
			sb.WriteRune(' ')
		}

		if entry.Err != nil {
			sb.WriteString(colorize(entry.Err, colorRed))
			sb.WriteRune(' ')
		}

		if entry.Duration > 0 {
			sb.WriteString(colorize(formatDuration(entry.Duration), colorMagenta))
			sb.WriteRune(' ')
		}

		sb.WriteRune('\n')
	}

	sb.WriteRune('\n')

	return sb.String()
}
