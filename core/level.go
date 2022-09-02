package core

import "fmt"

type Enabler interface {
	Enable(Level) bool
}

type Level uint8

func (l Level) String() string {
	str := ""
	switch l {
	case DebugLevel:
		str = "DEBUG"
	case InfoLevel:
		str = "INFO"
	case WarnLevel:
		str = "WARN"
	case ErrorLevel:
		str = "ERROR"
	case FatalLevel:
		str = "FATAL"
	}
	return fmt.Sprintf("%s", str)
}

const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)
