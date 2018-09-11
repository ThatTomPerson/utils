package null_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thattomperson/utils/null"
)

func TestString(t *testing.T) {
	t.Run("valid string", func(t *testing.T) {
		v := null.String("foo")

		assert.True(t, v.Valid)
		assert.Equal(t, "foo", v.String)
	})

	t.Run("empty string", func(t *testing.T) {
		v := null.String("")

		assert.True(t, v.Valid)
		assert.Equal(t, "", v.String)
	})

	t.Run("nil string", func(t *testing.T) {
		v := null.String(nil)

		assert.False(t, v.Valid)
		assert.Equal(t, "", v.String)
	})
}

func TestBool(t *testing.T) {
	t.Run("true bool", func(t *testing.T) {
		v := null.Bool(true)

		assert.True(t, v.Valid)
		assert.True(t, v.Bool)
	})

	t.Run("false bool", func(t *testing.T) {
		v := null.Bool(false)

		assert.True(t, v.Valid)
		assert.False(t, v.Bool)
	})

	t.Run("nil bool", func(t *testing.T) {
		v := null.Bool(nil)

		assert.False(t, v.Valid)
		assert.False(t, v.Bool)
	})

	t.Run("invalid bool", func(t *testing.T) {
		v := null.Bool("bar")

		assert.True(t, v.Valid)
		assert.False(t, v.Bool)
	})

	t.Run("int bool true", func(t *testing.T) {
		v := null.Bool(1)

		assert.True(t, v.Valid)
		assert.True(t, v.Bool)
	})

	t.Run("int bool false", func(t *testing.T) {
		v := null.Bool(0)

		assert.True(t, v.Valid)
		assert.False(t, v.Bool)
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
		{desc: "string", val: "5", expected: 5},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := null.Int(tC.val)

			assert.Equal(t, tC.expected, v.Int64)
			assert.Equal(t, tC.expected > 0, v.Valid)
		})
	}
}

func TestFloat(t *testing.T) {
	t.Run("float64", func(t *testing.T) {
		v := null.Float(float64(1.0))

		assert.Equal(t, float64(1.0), v.Float64)
		assert.True(t, v.Valid)
	})

	t.Run("float32", func(t *testing.T) {
		v := null.Float(float32(1.0))

		assert.Equal(t, float64(1.0), v.Float64)
		assert.True(t, v.Valid)
	})

	t.Run("nil", func(t *testing.T) {
		v := null.Float(nil)

		assert.Equal(t, float64(0), v.Float64)
		assert.False(t, v.Valid)
	})
}
