package core

import (
	"fmt"
	"strings"
)

type Caller struct {
	File string
	Line int
}

func NewCaller(f string, l int) *Caller {
	return &Caller{
		File: f,
		Line: l,
	}
}

func (c *Caller) Format() string {
	f := c.File
	idx := strings.LastIndexByte(c.File, '/')
	if idx != -1 {
		idx = strings.LastIndexByte(c.File[:idx], '/')
		if idx != -1 {
			f = c.File[idx+1:]
		}
	}
	return fmt.Sprintf("%s:%d", f, c.Line)
}
