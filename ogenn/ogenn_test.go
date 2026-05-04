package ogenn

import (
	"testing"

	"git.restpa.com/packages/golang/n.git"

	"github.com/stretchr/testify/assert"
)

// Test struct
// NewNilString returns new NilString with value set to v.
func NewNilString(v string) NilString {
	return NilString{
		Value: v,
	}
}

// NilString is nullable string.
type NilString struct {
	Value string
	Null  bool
}

// SetTo sets value to v.
func (o *NilString) SetTo(v string) {
	o.Null = false
	o.Value = v
}

// IsNull returns true if value is Null.
func (o NilString) IsNull() bool { return o.Null }

// SetToNull sets value to null.
func (o *NilString) SetToNull() {
	o.Null = true
	var v string
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o NilString) Get() (v string, ok bool) {
	if o.Null {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o NilString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

func TestTo_Valid(t *testing.T) {
	src := n.New("hello", true)
	dst := n.Null[string]{}

	To(src, &dst)

	assert.True(t, dst.IsValid())
	assert.Equal(t, "hello", dst.Value)
}

func TestTo_Null(t *testing.T) {
	src := n.New(42, false)
	dst := n.New(99, true)

	To(src, &dst)

	assert.True(t, dst.IsNull())
}

func TestTo_NilStringToNull(t *testing.T) {
	src := NewNilString("abc")
	dst := n.Null[string]{}

	To(src, &dst)

	assert.True(t, dst.IsValid())
	assert.Equal(t, "abc", dst.Value)
}

func TestTo_NilStringNull(t *testing.T) {
	src := NilString{Null: true}
	dst := n.New("fallback", true)

	To(src, &dst)

	assert.True(t, dst.IsNull())
}

func TestTo_Bool(t *testing.T) {
	src := n.New(true, true)
	dst := n.Null[bool]{}

	To(src, &dst)

	assert.True(t, dst.IsValid())
	assert.True(t, dst.Value)
}

func TestTo_Struct(t *testing.T) {
	type Item struct {
		Name string
	}

	src := n.New(Item{Name: "test"}, true)
	dst := n.Null[Item]{}

	To(src, &dst)

	assert.True(t, dst.IsValid())
	assert.Equal(t, "test", dst.Value.Name)
}

func TestTo_OverwriteValidWithNull(t *testing.T) {
	src := n.New(0, false)
	dst := n.New(100, true)

	To(src, &dst)

	assert.True(t, dst.IsNull())
	assert.Equal(t, 0, dst.Value)
}
