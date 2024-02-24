package common

import (
	"log"
	"os"
)

type EmptyWriter struct{}

// Write implements io.Writer.
func (EmptyWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

var (
	LogError   = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogWarning = log.New(os.Stderr, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogInfo    = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogDebug   = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	LogEmpty   = log.New(EmptyWriter{}, "", 0)

	LogLevel = Error | Warn | Info | Debug
)

type Level int

const (
	Error Level = 1 << iota
	Warn  Level = 1 << iota
	Info  Level = 1 << iota
	Debug Level = 1 << iota
)

func Logger(level Level) *log.Logger {
	if LogLevel&level != 0 {
		switch level {
		case Warn:
			return LogWarning
		case Info:
			return LogInfo
		case Debug:
			return LogDebug
		case Error:
			return LogError
		}
	}
	return LogEmpty
}

func SetLogLevel(level Level) {
	LogLevel = 0
	switch level {
	case Debug:
		LogLevel |= Debug
		fallthrough
	case Info:
		LogLevel |= Info
		fallthrough
	case Warn:
		LogLevel |= Warn
		fallthrough
	case Error:
		LogLevel |= Error
	}
}
