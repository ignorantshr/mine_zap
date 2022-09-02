package multierr

import (
	"bytes"
	"io"
	"sync"
)

var (
	_multiErrSpiltor = []byte("; ")
)

var _bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

type errorGroup interface {
	Errors() []error
}

var _ errorGroup = (*multiErr)(nil)

type multiErr struct {
	errs []error
}

func (mmr *multiErr) Errors() []error {
	if mmr == nil {
		return nil
	}
	return mmr.errs
}

func (mmr *multiErr) Error() string {
	if mmr == nil {
		return ""
	}

	buff := _bufferPool.Get().(*bytes.Buffer)
	buff.Reset()

	mmr.writeSingleLine(buff)

	_bufferPool.Put(buff)
	return buff.String()
}

func (mmr *multiErr) writeSingleLine(w io.Writer) {
	first := true
	for _, err := range mmr.errs {
		if first {
			first = false
		} else {
			w.Write(_multiErrSpiltor)
		}
		io.WriteString(w, err.Error())
	}
}

type errInspector struct {
	counts  int
	indexes []int
}

func inspect(errs []error) errInspector {
	idx := make([]int, 0)
	count := 0
	for i, err := range errs {
		if err != nil {
			idx = append(idx, i)
			count++
		}
	}
	return errInspector{
		counts:  count,
		indexes: idx,
	}
}

func fromSlice(errors []error) error {
	res := inspect(errors)
	if res.counts == 0 {
		return nil
	}

	errs := make([]error, 0)
	for _, i := range res.indexes {
		if nested, ok := errors[i].(*multiErr); !ok {
			errs = append(errs, errors[i])
		} else {
			errs = append(errs, nested.Errors()...)
		}
	}
	return &multiErr{
		errs: errs,
	}
}

func Combine(errs ...error) error {
	return fromSlice(errs)
}
