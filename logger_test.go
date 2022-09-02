package mine_zap

import (
	"log"
	"mine_zap/core"
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

	ol := New(core.NewCore(core.InfoLevel, core.NewJsonEncoder()))
	ol.Info("origin", AnyType("A", "a~"), Error("fake error occurs"))
}
