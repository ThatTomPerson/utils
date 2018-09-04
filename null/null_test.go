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
