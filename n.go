package n

import (
	"bytes"
	"encoding/json"
)

type Null[T any] struct {
	Value T
	Valid bool
}

func New[T any](value T, valid bool) Null[T] {
	return Null[T]{
		Value: value,
		Valid: valid,
	}
}

func (n Null[T]) IsValid() bool { return n.Valid }

func (n Null[T]) IsNull() bool { return !n.Valid }

// SetTo sets value to v.
func (n *Null[T]) SetTo(v T) {
	n.Valid = true
	n.Value = v
}

// SetNull sets value to null.
func (n *Null[T]) SetToNull() {
	n.Valid = false
	var v T
	n.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (n Null[T]) Get() (v T, ok bool) {
	if n.IsNull() {
		return v, false
	}
	return n.Value, true
}

// Or returns value if set, or given parameter if does not.
func (n Null[T]) Or(d T) T {
	if v, ok := n.Get(); ok {
		return v
	}
	return d
}

func (n Null[T]) MarshalJSON() ([]byte, error) {
	if n.IsNull() {
		return []byte("null"), nil
	}

	return json.Marshal(n.Value)
}

func (n *Null[T]) UnmarshalJSON(data []byte) error {
	if data == nil || len(data) == 0 {
		n.SetToNull()
		return nil
	}

	if bytes.Equal(data, []byte("null")) {
		n.SetToNull()
		return nil
	}
	n.Valid = true
	return json.Unmarshal(data, &n.Value)
}
