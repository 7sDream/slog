package slog

import (
	"fmt"
	"log"
	"os"
)

// Logger is a util class to print log in different level like DEBUG.
// It has 4 level: DEBUG, INFO, WARN, ERROR,
// and 2 exception logger: PANIC and FATAL
type Logger struct {
	Prefix string
	dl     *log.Logger
	il     *log.Logger
	wl     *log.Logger
	el     *log.Logger
	pl     *log.Logger
	fl     *log.Logger
}

// Logger level const vars
const (
	LevelDebug int = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	l = NewLogger("[Logger]")
	logFlag = log.Ldate | log.Ltime | log.Lmicroseconds

	currentLevel = LevelDebug
)

// NewLogger create a new top level logger based on prefix.
func NewLogger(prefix string) *Logger {
	return &Logger{
		Prefix: prefix,
		dl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "DEBUG", prefix), logFlag),
		il:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "INFO", prefix), logFlag),
		wl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "WARN", prefix), logFlag),
		el:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "ERROR", prefix), logFlag),
		pl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "PANIC", prefix), logFlag),
		fl:     log.New(os.Stdout, fmt.Sprintf("[%s] %s ", "FATAL", prefix), logFlag),
	}
}

// Debug print log message as DEBUG level.
// if you set log level higher than LevelDebug, no message will be print.
func (logger *Logger) Debug(data ... interface{}) {
	if currentLevel <= LevelDebug {
		logger.dl.Println(data...)
	}
}

// Info print log message as INFO level.
// If you set log level higher than LevelInfo, no message will be print.
func (logger *Logger) Info(data ... interface{}) {
	if currentLevel <= LevelInfo {
		logger.il.Println(data...)
	}
}

// Warn print log message as WARN level.
// If you set log level higher than LevelWarn, no message will be print.
func (logger *Logger) Warn(data ... interface{}) {
	if currentLevel <= LevelWarn {
		logger.wl.Println(data...)
	}
}

// Error print log message as ERROR level.
// This function do not create panic or fatal, it just print error message.
// If you want get a runtime panic or fatal, use Logger.Panic or Logger.Fatal instand.
func (logger *Logger) Error(data ... interface{}) {
	if currentLevel <= LevelError {
		logger.el.Println(data...)
	}
}

// Panic print log message, and create a panic use the message.
func (logger *Logger) Panic(data ... interface{}) {
	logger.pl.Panicln(data...)
}

// Fatal print log message, and create a fatal use the message.
func (logger *Logger) Fatal(data ... interface{}) {
	logger.fl.Fatalln(data...)
}

// SubLogger create a new logger based on the logger.
// Prefix string of new logger will be concat of old and provided argument.
func (logger *Logger) SubLogger(prefix string) (subLogger *Logger) {
	subLogger = NewLogger(fmt.Sprintf("%s %s", logger.Prefix, prefix))
	return subLogger
}

// SetLevel set the minimum level that message will be print out
func SetLevel(level int) {
	if LevelDebug <= level && level <= LevelError {
		currentLevel = level
	} else {
		l.Error("Set logger level", level, "failed, accepted range is", LevelDebug, "to", LevelError)
	}
}
