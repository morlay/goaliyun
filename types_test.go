package goaliyun

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	tt := assert.New(t)

	{
		v := String("")
		err := v.UnmarshalJSON([]byte(`"1"`))
		tt.NoError(err)
		tt.Equal(string(v), "1")
	}

	{
		v := String("")
		err := v.UnmarshalJSON([]byte(`1`))
		tt.NoError(err)
		tt.Equal(string(v), "1")
	}
}

func TestInteger(t *testing.T) {
	tt := assert.New(t)

	{
		v := Integer(0)
		err := v.UnmarshalJSON([]byte(`"1"`))
		tt.NoError(err)
		tt.Equal(int(v), int(1))
	}

	{
		v := Integer(0)
		err := v.UnmarshalJSON([]byte(`1`))
		tt.NoError(err)
		tt.Equal(int(v), int(1))
	}
}

func TestFloat(t *testing.T) {
	tt := assert.New(t)

	{
		v := Float(0)
		err := v.UnmarshalJSON([]byte(`"1.1"`))
		tt.NoError(err)
		tt.Equal(float32(v), float32(1.1))
	}

	{
		v := Float(0)
		err := v.UnmarshalJSON([]byte(`1.1`))
		tt.NoError(err)
		tt.Equal(float32(v), float32(1.1))
	}
}
