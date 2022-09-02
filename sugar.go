package mine_zap

import (
	"fmt"
	"mine_zap/core"
)

type SugaredLogger struct {
	base *Logger
}

func NewSugaredLogger(level core.Level, options ...Option) *SugaredLogger {
	c := core.NewCore(level, core.NewJsonEncoder())
	base := New(c, options...)
	base.callerSkip = 2 // 在基础的 skip 上面额外跳过的栈数
	return &SugaredLogger{
		base: base,
	}
}

func (l *SugaredLogger) Desugar() *Logger {
	log := l.base.clone()
	log.callerSkip -= 2
	return log
}

func (l *SugaredLogger) Info(args ...interface{}) {
	l.log(core.InfoLevel, "", args)
}

func (l *SugaredLogger) Infof(template string, args ...interface{}) {
	l.log(core.InfoLevel, template, args)
}

func (l *SugaredLogger) log(level core.Level, template string, args []interface{}) {
	if !l.base.core.Enable(level) {
		return
	}

	m := template
	if template != "" && len(args) > 0 {
		m = fmt.Sprintf(template, args...)
	} else if template == "" && len(args) > 0 {
		m = fmt.Sprint(args...)
	}

	ce := l.base.Check(level, m)
	ce.Write()
}
