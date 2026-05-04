package n

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestEncoding(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			name: "valid int",
			in:   Null[int]{Value: 1, Valid: true},
			want: "1",
		},
		{
			name: "null int",
			in:   Null[int]{Value: 1, Valid: false},
			want: "null",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			at := assert.New(t)
			data, err := json.Marshal(tt.in)
			at.Nil(err, "error should be nil")
			out := string(data)
			at.Equal(out, tt.want, "output and want objects should be equal")
		})
	}
}

func TestDecodingPrimitive(t *testing.T) {
	tests := []struct {
		name string
		in   []byte
		want Null[any]
	}{
		{
			name: "null",
			in:   []byte("null"),
			want: Null[any]{Value: nil, Valid: false},
		},
		{
			name: "valid_int",
			in:   []byte("2"),
			want: Null[any]{Value: 2, Valid: true},
		},
		{
			name: "valid_float",
			in:   []byte("2.123"),
			want: Null[any]{Value: 2.123, Valid: true},
		},
		{
			name: "valid_bool",
			in:   []byte("true"),
			want: Null[any]{
				Value: true,
				Valid: true,
			},
		},
	}
	for _, tt := range tests {
		at := assert.New(t)
		t.Run(tt.name, func(t *testing.T) {
			out := tt.want
			err := json.Unmarshal(tt.in, &out)
			at.Nil(err, "error should be nil")
			at.Equal(tt.want.IsValid(), out.IsValid(), "output and want validity should be equal")
			at.EqualValues(tt.want.Value, out.Value, tt.name)
		})
	}
}

func TestDecodingStruct(t *testing.T) {
	type Item struct {
		Name string
		Time time.Time
	}

	in := []byte("{\"Name\":\"test1\",\"Time\":\"2000-07-19T00:00:00Z\"}")
	want := Null[Item]{
		Value: Item{
			Name: "test1",
			Time: time.Date(2000, time.July, 19, 0, 0, 0, 0, time.UTC),
		},
		Valid: true,
	}
	var out Null[Item]
	at := assert.New(t)
	err := json.Unmarshal(in, &out)
	at.Nil(err, "error should be nil")
	at.Equal(want.IsValid(), out.IsValid(), "output and want validity should be equal")
	at.EqualValues(want.Value, out.Value, "output and want objects should be equal")
}

func TestDecodingTime(t *testing.T) {
	in := []byte("\"2000-07-19T00:00:00Z\"")
	want := Null[time.Time]{
		Value: time.Date(2000, time.July, 19, 0, 0, 0, 0, time.UTC),
		Valid: true,
	}
	var out Null[time.Time]
	at := assert.New(t)
	err := json.Unmarshal(in, &out)
	at.Nil(err, "error should be nil")
	at.Equal(want.IsValid(), out.IsValid(), "output and want validity should be equal")
	at.EqualValues(want.Value, out.Value, "output and want objects should be equal")
}
