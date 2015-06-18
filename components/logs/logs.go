package logs

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

func logger(level string, c color.Attribute) func(format interface{}, v ...interface{}) {
	colorPrint := color.New(c).SprintFunc()
	return func(format interface{}, v ...interface{}) {
		switch format.(type) {
		case string:
			log.Printf("%s %s\n", colorPrint(level), fmt.Sprintf(format.(string), v...))
		case error:
			log.Printf("%s %s\n", colorPrint(level), fmt.Sprint(format.(error)))
		default:
			log.Printf("%s %v\n", colorPrint(level), format)
		}
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
