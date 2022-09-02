package mine_zap

import "mine_zap/core"

type Field = core.Field

func AnyType(key string, val interface{}) Field {
	return Field{
		Key:       key,
		Interface: val,
	}
}
