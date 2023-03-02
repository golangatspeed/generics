// Concrete testing approach
package generic

import "testing"

func TestAddAnyInt8(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		x    int8
		y    int8
		want int8
	}{
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

func TestAddAnyString(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		x    string
		y    string
		want string
	}{
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
