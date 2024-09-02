package gox

import (
	"encoding/json"
	"encoding/xml"
)

// Slice 切片，既可以兼容单个值也可以兼容数组
type Slice[T any] []T

// NewSlice 快速创建数组
func NewSlice[T any](items ...T) Slice[T] {
	return items
}

func (s *Slice[T]) Length() int {
	return len(*s)
}

func (s *Slice[T]) Capacity() int {
	return cap(*s)
}

func (s *Slice[T]) Clone() (t Slice[T]) {
	t = make(Slice[T], len(*s))
	copy(t, *s)

	return
}

func (s *Slice[T]) UnmarshalJSON(bytes []byte) (err error) {
	t := new(T)
	if ue := json.Unmarshal(bytes, t); nil != ue {
		err = json.Unmarshal(bytes, s)
	} else {
		*s = []T{*t}
	}

	return
}

func (s *Slice[T]) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) (err error) {
	t := new(T)
	if ue := decoder.DecodeElement(t, &start); nil != ue {
		err = decoder.DecodeElement(s, &start)
	} else {
		*s = []T{*t}
	}

	return
}
