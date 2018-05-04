// Package logg defines the log handling used by all blog-004-errlog software.
package logg

import (
	"io"
	"log"

	"github.com/snassr/blog-0004-errlog/errors"
)

// Kind defines the kind of logging.
type Kind int

// Kinds of logging
//
// The values of the log kinds are common between both
// clients and servers (order is important, list is append-only).
// Important update the count.
const (
	Trace Kind = iota + 1 // Track the code operations (Logical Line Posting).
	Info                  // Useful program information (start/stop, config, development envrionment).
	Warn                  // Potiential causes for recoverable application oddities (recovery, retry, missing).
	Err                   // An error which is fatal to the operation but not program.
	Fatal                 // An error that is fatal to the program (data impact situations).
)

// Logger is a log management object.
type Logger struct {
	logTrace *log.Logger
	logInfo  *log.Logger
	logWarn  *log.Logger
	logErr   *log.Logger
	logFatal *log.Logger
}

// Init intializes loggers for an application runtime and
// returns an instance of Logger.
// Each log level can have multiple writers (stdout, database, file...).
func Init(trace, info, warn, err, fatal []io.Writer) *Logger {
	mwtrace := io.MultiWriter(trace...)
	mwinfo := io.MultiWriter(info...)
	mwwarn := io.MultiWriter(warn...)
	mwerr := io.MultiWriter(err...)
	mwfatal := io.MultiWriter(fatal...)

	return &Logger{
		logTrace: log.New(mwtrace, "[TRACE]: ", log.Ldate|log.Ltime|log.Lmicroseconds),
		logInfo:  log.New(mwinfo, "[INFO]: ", log.Ldate|log.Ltime|log.Lmicroseconds),
		logWarn:  log.New(mwwarn, "[WARN]: ", log.Ldate|log.Ltime|log.Lmicroseconds),
		logErr:   log.New(mwerr, "[ERROR]: ", log.Ldate|log.Ltime|log.Lmicroseconds),
		logFatal: log.New(mwfatal, "[FATAL]: ", log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

// Trace will log an error at the Trace level.
func (l *Logger) Trace(err errors.Error) {
	l.logTrace.Printf(string(err.JSON()))
}

// Info will log an error at the Info level.
func (l *Logger) Info(err errors.Error) {
	l.logInfo.Printf(string(err.JSON()))
}

// Warn will log an error at the Trace level.
func (l *Logger) Warn(err errors.Error) {
	l.logWarn.Printf(string(err.JSON()))
}

// Err will log an error at the Err level.
func (l *Logger) Err(err errors.Error) {
	l.logErr.Printf(string(err.JSON()))
}

// Fatal will log an error at the Fatal level.
func (l *Logger) Fatal(err errors.Error) {
	l.logFatal.Printf(string(err.JSON()))
}
