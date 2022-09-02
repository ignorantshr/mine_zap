package core

type ObjectEncoder interface {
	AddString(key string, val string)
	AddBool(key string, val bool)
	AddInt(key string, val int)
	AddInt64(key string, val int64)
	AddAny(key string, val interface{})
}

type Encoder interface {
	ObjectEncoder
	Encode(Entry) ([]byte, error)
}
