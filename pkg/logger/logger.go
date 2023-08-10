// Package logger provides a logging package for writing messages to various 
// outputs.
package logger

import (
	"fmt"
	"log"
	"os"
)

var infoLogger *log.Logger
var warnLogger *log.Logger
var errLogger *log.Logger
var fatalLogger *log.Logger

// Init initializes four loggers with different severity levels and sets up
// their respective output files. It returns an error if any of the log files
// fail to be created.
func Init(infoPath, warningPath, errorPath, fatalPath string) error {
	logConfigs := [...]struct {
		logger **log.Logger
		path   string
		prefix string
		flags  int
	}{
		{&infoLogger, infoPath, "INFO: ", log.LstdFlags},
		{&warnLogger, warningPath, "WARN: ", log.LstdFlags | log.Llongfile},
		{&errLogger, errorPath, "ERROR: ", log.LstdFlags | log.Llongfile},
		{&fatalLogger, fatalPath, "FATAL: ", log.LstdFlags | log.Llongfile},
	}

	fileFlags := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	for _, logConfig := range logConfigs {
		file, err := os.OpenFile(logConfig.path, fileFlags, 0755)
		if err != nil {
			return fmt.Errorf("failed to create log file for %s: %w", logConfig.path, err)
		}
		*logConfig.logger = log.New(file, logConfig.prefix, logConfig.flags)
	}
	return nil
}

// Info writes messages to the infoLogger at the INFO severity level.
func Info(format string, v ...any) {
	output(infoLogger, format, v...)
}

// Warn writes messages to the warnLogger at the WARN severity level.
func Warn(format string, v ...any) {
	output(warnLogger, format, v...)
}

// Warn writes messages to the errLogger at the ERROR severity level.
func Err(format string, v ...any) {
	output(errLogger, format, v...)
}

// Fatal writes messages to the fatalLogger at the FATAL severity level and
// exits the program with a status code of 1.
func Fatal(format string, v ...any) {
	output(fatalLogger, format, v...)
	os.Exit(1)
}

// output is used internally to format and write log messages to the specified
// logger and also to the console. If the logger is not initialized, a panic
// occurs.
func output(logger *log.Logger, format string, v ...any) {
	if logger == nil {
		panic("logger is not initialized")
	}
	logger.Output(2, fmt.Sprintf(format, v...))

	oldWriter := logger.Writer()
	if oldWriter != os.Stderr {
		logger.SetOutput(os.Stderr)
		logger.Output(2, fmt.Sprintf(format, v...))
		logger.SetOutput(oldWriter)
	}
}
