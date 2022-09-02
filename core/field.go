package core

type Field struct {
	Key string
	//Integer   int64
	//String    string
	Interface interface{}
}

func NewField(key string, val interface{}) Field {
	return Field{
		Key: key,
		//Integer:   0,
		//String:    "",
		Interface: val,
	}
}
