package logger

import "fmt"

var AlpacaLoggerEnabled = true
var AlpacaLoggerExtraEnabled = false

func Log(format string, v ...any) {
	if !AlpacaLoggerEnabled {
		return
	}
	fmt.Printf(format, v...)
}

func LogExtra(format string, v ...any) {
	if !AlpacaLoggerExtraEnabled {
		return
	}
	fmt.Printf(format, v...)
}

func LogError(format string, v ...any) {
	fmt.Printf(format, v...)
}
