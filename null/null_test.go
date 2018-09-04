package null_test

import (
	"fmt"
	"testing"

	"github.com/ThatTomPerson/utils/null"
	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		v := null.String("foo")

		assert.True(t, v.Valid, "string is not valid")
		assert.Equal(t, v.String, "foo", fmt.Sprintf("expected foo but got %s", v.String))
	})

	t.Run("valid string", func(t *testing.T) {
		v := null.String("")

		assert.False(t, v.Valid, "string is valid")
		assert.Equal(t, v.String, "", fmt.Sprintf("expected foo but got %s", v.String))
	})
}

func TestBool(t *testing.T) {
	t.Run("true bool", func(t *testing.T) {
		v := null.Bool(true)

		assert.True(t, v.Valid, "valid is true")
		assert.Equal(t, v.Bool, true, "expected true but got false")
	})

	t.Run("false bool", func(t *testing.T) {
		v := null.Bool(false)

		assert.True(t, v.Valid, "valid is true")
		assert.Equal(t, v.Bool, false, "expected false but got true")
	})

	t.Run("nil bool", func(t *testing.T) {
		v := null.Bool(nil)

		assert.False(t, v.Valid, "valid is true")
		assert.Equal(t, v.Bool, false, "expected false but got true")
	})

	t.Run("invalid bool", func(t *testing.T) {
		v := null.Bool("bar")

		assert.False(t, v.Valid, "valid is true")
		assert.Equal(t, v.Bool, false, "expected false but got true")
	})
}

func TestInt(t *testing.T) {
	testCases := []struct {
		desc     string
		val      interface{}
		expected int64
	}{
		{desc: "int", val: int(10), expected: 10},
		{desc: "int8", val: int8(10), expected: 10},
		{desc: "int16", val: int16(10), expected: 10},
		{desc: "int32", val: int32(10), expected: 10},
		{desc: "int64", val: int64(10), expected: 10},
		{desc: "uint", val: uint(10), expected: 10},
		{desc: "uint8", val: uint8(10), expected: 10},
		{desc: "uint16", val: uint16(10), expected: 10},
		{desc: "uint32", val: uint32(10), expected: 10},
		{desc: "uint64", val: uint64(10), expected: 10},
		{desc: "nil", val: nil, expected: 0},
		{desc: "string", val: "test", expected: 0},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := null.Int(tC.val)

			assert.Equal(t, v.Int64, tC.expected, fmt.Sprintf("expected %d but got %d", tC.expected, v.Int64))
			assert.Equal(t, v.Valid, tC.expected > 0, fmt.Sprintf("expected %v but got %v", tC.expected > 0, v.Valid))
		})
	}
}

func TestFloat(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		v := null.Float(float64(1.0))

		assert.Equal(t, float64(1.0), v.Float64, fmt.Sprintf("expected %f but got %f", float64(1.0), v.Float64))
		assert.True(t, v.Valid, "expected true but got false")
	})

	t.Run("float32", func(t *testing.T) {
		v := null.Float(float32(1.0))

		assert.Equal(t, float64(1.0), v.Float64, fmt.Sprintf("expected %f but got %f", float64(1.0), v.Float64))
		assert.True(t, v.Valid, "expected true but got false")
	})

	t.Run("nil", func(t *testing.T) {
		v := null.Float(nil)

		assert.Equal(t, float64(0), v.Float64, fmt.Sprintf("expected %f but got %f", float64(0), v.Float64))
		assert.False(t, v.Valid, "expected false but got true")
	})
}
