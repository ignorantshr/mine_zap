package core

import (
	"io"
	"os"
)

type Core interface {
	Enabler

	With([]Field) Core

	Check(Entry, *CheckedEntry) *CheckedEntry

	Write(Entry, []Field) error
}

var (
	defaultWriter = os.Stderr
)

type ioCore struct {
	level Level
	out   io.Writer
	enc   Encoder
}

func NewCore(level Level, encoder Encoder) Core {
	return &ioCore{
		level: level,
		out:   defaultWriter,
		enc:   encoder,
	}
}

func (c *ioCore) Enable(level Level) bool {
	return c.level <= level
}

func (c *ioCore) With(fields []Field) Core {
	for _, f := range fields {
		c.enc.AddAny(f.Key, f.Interface)
	}
	return c
}

func (c *ioCore) Check(entry Entry, checkedEnt *CheckedEntry) *CheckedEntry {
	if !c.Enable(entry.Level) {
		return nil
	}
	return checkedEnt.AddCore(entry, c)
}

func (c *ioCore) Write(entry Entry, fields []Field) error {
	res, err := c.enc.Encode(entry, fields)
	if err != nil {
		return err
	}
	_, err = c.out.Write(append(res, '\r', '\n'))
	return err
}
