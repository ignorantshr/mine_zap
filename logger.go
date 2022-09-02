package mine_zap

import (
	"mine_zap/core"
)

type Logger struct {
	core core.Core

	callerSkip int
}

func New(core core.Core, options ...Option) *Logger {
	if core == nil {
		return nil
	}
	log := &Logger{
		core: core,
	}
	return log.WithOptions(options...)
}

func (l *Logger) Sugar() *SugaredLogger {
	base := l.clone()
	base.callerSkip += 1
	return &SugaredLogger{base: base}
}

// WithOptions 先 clone 保证线程安全
func (l *Logger) WithOptions(options ...Option) *Logger {
	log := l.clone()
	for _, opt := range options {
		opt.apply(log)
	}
	return log
}

func (l *Logger) With(fields ...Field) *Logger {
	log := l.clone()
	log.core = log.core.With(fields)
	return log
}

func (l *Logger) Debug(msg string) {
	if ce := l.check(core.DebugLevel, msg); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Info(msg string) {
	if ce := l.check(core.InfoLevel, msg); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Warn(msg string) {
	if ce := l.check(core.WarnLevel, msg); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Error(msg string) {
	if ce := l.check(core.ErrorLevel, msg); ce != nil {
		ce.Write()
	}
}

// Check 再封装一层是为了能确定 callerSkip 的基数
func (l *Logger) Check(level core.Level, msg string) *core.CheckedEntry {
	return l.check(level, msg)
}

func (l *Logger) check(level core.Level, msg string) *core.CheckedEntry {
	const callerSkipOffset = 3
	ent := core.NewEntry(level, msg, l.callerSkip+callerSkipOffset)
	return l.core.Check(*ent, nil)
}

func (l *Logger) clone() *Logger {
	log := *l
	return &log
}
