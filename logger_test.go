package mine_zap

import (
	"errors"
	"log"
	"mine_zap/core"
	"mine_zap/multierr"
	"testing"
)

func TestNewLogger(t *testing.T) {
	l := NewSugaredLogger(core.InfoLevel,
		Fields(AnyType("C", 1), AnyType("B", "b")),
		Hooks(
			func(entry core.Entry) error {
				log.Println("h1")
				return nil
			},
			func(entry core.Entry) error {
				log.Println("h2")
				return nil
			},
		),
		Hooks(func(entry core.Entry) error {
			log.Println("h3")
			return nil
		}),
	)
	l.Infof("hh: %s", "sd")
	logger := l.Desugar()
	logger.Info("haha")
	ol := New(core.NewCore(core.InfoLevel))
	ol.Info("origin")

	errGroup := multierr.Combine(errors.New("a\t"), nil)
	errGroup = multierr.Combine(errGroup, errors.New("c\n"))
	ol.Info(errGroup.Error())
}
