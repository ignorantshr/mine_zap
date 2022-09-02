package core

import (
	"encoding/json"
	"strconv"
)

type jsonEncoder struct {
	fields map[string]interface{}
}

func NewJsonEncoder() *jsonEncoder {
	return &jsonEncoder{
		fields: make(map[string]interface{}),
	}
}

func (je *jsonEncoder) AddString(key string, val string) {
	je.AddAny(key, val)
}

func (je *jsonEncoder) AddBool(key string, val bool) {
	je.AddAny(key, val)
}

func (je *jsonEncoder) AddInt(key string, val int) {
	je.AddAny(key, val)
}

func (je *jsonEncoder) AddInt64(key string, val int64) {
	je.AddAny(key, val)
}

func (je *jsonEncoder) AddAny(key string, val interface{}) {
	je.fields[key] = val
}

func (je *jsonEncoder) Encode(entry Entry, fields []Field) ([]byte, error) {
	header := je.encodeHeader(entry)

	for _, field := range fields {
		je.AddAny(field.Key, field.Interface)
	}
	if len(je.fields) != 0 {
		extra, err := json.Marshal(je.fields)
		if err != nil {
			return nil, err
		}
		header = append(header[:len(header)-1], ',')
		return append(header, extra[1:]...), nil
	}

	return header, nil
}

func (je *jsonEncoder) encodeHeader(entry Entry) []byte {
	res := []byte{'{'}
	formatHeader := func(k, v string) {
		res = strconv.AppendQuote(res, k)
		res = append(res, ':')
		res = strconv.AppendQuote(res, v)
	}
	formatHeader(LevelStr, entry.Level.String())
	res = append(res, ',')
	formatHeader(TimeStr, entry.Time.String())
	res = append(res, ',')
	formatHeader(CallerStr, entry.Caller.Format())
	res = append(res, ',')
	formatHeader(MessageStr, entry.Message)
	res = append(res, '}')
	return res
}
