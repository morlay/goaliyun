package goaliyun

import (
	"bytes"
	"strconv"
)

type String string

func (v String) String() string {
	return string(v)
}

func (v *String) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		data = bytes.Trim(data, `"`)
	}
	*v = String(data)
	return nil
}

type Integer int64

func (v Integer) Int() int {
	return int(v)
}

func (v Integer) Int64() int64 {
	return int64(v)
}

func (v *Integer) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		data = bytes.Trim(data, `"`)
	}
	i, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		return err
	}
	*v = Integer(i)
	return nil
}

type Float float64

func (v Float) Float32() float32 {
	return float32(v)
}

func (v Float) Float64() float64 {
	return float64(v)
}

func (v *Float) UnmarshalJSON(data []byte) error {
	if len(data) > 0 && data[0] == '"' {
		data = bytes.Trim(data, `"`)
	}
	f, err := strconv.ParseFloat(string(data), 32)
	if err != nil {
		return err
	}
	*v = Float(f)
	return nil
}
