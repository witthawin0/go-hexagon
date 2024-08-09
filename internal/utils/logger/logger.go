// internal/utils/logger/logger.go
package logger

import (
	"log"
	"os"
)

var (
	infoLogger    *log.Logger
	warningLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Info logs informational messages
func Info(v ...interface{}) {
	infoLogger.Println(v...)
}

// Warning logs warning messages
func Warning(v ...interface{}) {
	warningLogger.Println(v...)
}

// Error logs error messages
func Error(v ...interface{}) {
	errorLogger.Println(v...)
}

// Fatal logs error messages and then calls os.Exit(1)
func Fatal(v ...interface{}) {
	errorLogger.Fatal(v...)
}
