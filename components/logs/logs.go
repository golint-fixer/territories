package logs

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func logger(level string, c color.Attribute) func(format string, v ...interface{}) {
	colorPrint := color.New(c).SprintFunc()
	return func(format string, v ...interface{}) {
		log.Printf("%s %s\n", colorPrint(level), fmt.Sprintf(format, v...))
	}
}

// Debug logger
var LogDebug = logger("DEBU", color.FgCyan)

// Info logger
var LogInfo = logger("INFO", color.Reset)

// Notice logger
var LogNotice = logger("NOTI", color.FgGreen)

// Warning logger
var LogWarning = logger("WARN", color.FgYellow)

// Error logger
var LogError = logger("ERRO", color.FgRed)

// Critical logger
var LogCritical = logger("CRIT", color.FgMagenta)
