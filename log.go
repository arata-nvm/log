package log

import (
	"fmt"
	"log"

	"github.com/fatih/color"
)

var prefix = map[string]string{
	"status":    fmt.Sprintf("[%s]", color.MagentaString("x")),
	"success":   fmt.Sprintf("[%s]", color.GreenString("+")),
	"failure":   fmt.Sprintf("[%s]", color.RedString("-")),
	"debug":     fmt.Sprintf("[%s]", color.RedString("DEBUG")),
	"info":      fmt.Sprintf("[%s]", color.BlueString("*")),
	"warning":   fmt.Sprintf("[%s]", color.YellowString("!")),
	"error":     fmt.Sprintf("[%s]", color.RedString("ERROR")),
	"exception": fmt.Sprintf("[%s]", color.RedString("ERROR")),
	"critical":  fmt.Sprintf("[%s]", color.RedString("CRITICAL")),
}

func Status(format string, v ...interface{}) {
	log.Printf(prefix["status"]+format, v)
}

func Success(format string, v ...interface{}) {
	log.Printf(prefix["success"]+format, v)
}

func Failure(format string, v ...interface{}) {
	log.Printf(prefix["failure"]+format, v)
}

func Debug(format string, v ...interface{}) {
	log.Printf(prefix["debug"]+format, v)
}

func Info(format string, v ...interface{}) {
	log.Printf(prefix["info"]+format, v)
}

func Warning(format string, v ...interface{}) {
	log.Printf(prefix["warning"]+format, v)
}

func Error(format string, v ...interface{}) {
	log.Printf(prefix["error"]+format, v)
}

func Exeption(format string, v ...interface{}) {
	log.Printf(prefix["exeption"]+format, v)
}

func Critical(format string, v ...interface{}) {
	log.Printf(prefix["critical"]+format, v)
}
