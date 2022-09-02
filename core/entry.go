package core

import (
	"runtime"
	"sync"
	"time"
)

type Entry struct {
	Level
	Time time.Time
	*Caller
	Message string
}

func NewEntry(level Level, message string, callerSkip int) *Entry {
	_, f, l, _ := runtime.Caller(callerSkip)
	return &Entry{
		Level:   level,
		Time:    time.Now(),
		Caller:  NewCaller(f, l),
		Message: message,
	}
}

var (
	_cePool = sync.Pool{New: func() interface{} {
		return &CheckedEntry{
			cores: make([]Core, 1),
		}
	}}
)

type CheckedEntry struct {
	Entry
	cores []Core
}

func (ce *CheckedEntry) reset() {
	ce.Entry = Entry{}
	for i := range ce.cores {
		ce.cores[i] = nil
	}
	ce.cores = ce.cores[:0]
}

func (ce *CheckedEntry) AddCore(ent Entry, core Core) *CheckedEntry {
	if ce == nil {
		ce = getCheckedEntry()
		ce.Entry = ent
	}
	ce.cores = append(ce.cores, core)
	return ce
}

func (ce *CheckedEntry) Write() {
	if ce == nil {
		return
	}
	for _, c := range ce.cores {
		_ = c.Write(ce.Entry)
	}
	putCheckedEntry(ce)
}

func getCheckedEntry() *CheckedEntry {
	ce := _cePool.Get().(*CheckedEntry)
	ce.reset()
	return ce
}

func putCheckedEntry(ce *CheckedEntry) {
	if ce == nil {
		return
	}
	_cePool.Put(ce)
}
