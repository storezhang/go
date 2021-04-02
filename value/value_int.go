package value

import (
	`github.com/storezhang/gox`
)

var _ gox.Value = (*IntValue)(nil)

// IntValue 整形值
type IntValue struct {
	value int
}

// Int 整形值
func Int(value int) IntValue {
	return IntValue{
		value: value,
	}
}

func (iv *IntValue) Value() interface{} {
	return iv.value
}