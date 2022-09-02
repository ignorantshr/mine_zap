package mine_zap

import "mine_zap/core"

type Option interface {
	apply(*Logger)
}

type OptionFunc func(*Logger)

func (fn OptionFunc) apply(l *Logger) {
	fn(l)
}

func Fields(fields ...Field) Option {
	return OptionFunc(func(l *Logger) {
		l.core.With(fields)
	})
}

func Hooks(hooks ...core.HookFunc) Option {
	return OptionFunc(func(l *Logger) {
		l.core = core.RegisterHooks(l.core, hooks...)
	})
}
