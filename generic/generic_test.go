// Generic testing approach
package generic

import "testing"

type testCase[T Addable] struct {
	name string
	x    T
	y    T
	want T
}

func TestAddAnyGenInt8(t *testing.T) {
	t.Parallel()
	tests := []testCase[int8]{
		{
			"first test",
			2,
			2,
			4,
		},
		{
			"second test",
			0,
			-2,
			-2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAny(tt.x, tt.y); got != tt.want {
				t.Errorf("AddAny() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddAnyGenString(t *testing.T) {
	t.Parallel()
	tests := []testCase[string]{
		{
			"first test string",
			"hello",
			" world",
			"hello world",
		},
		{
			"second test string",
			"My name is ",
			"Joe Blogs",
			"My name is Joe Blogs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAny(tt.x, tt.y); got != tt.want {
				t.Errorf("AddAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
