package generic

import "testing"

/*
func TestAddAny(t *testing.T) {

	type testCase[T Addable] struct {
		name string
		x    T
		y    T
		want T
	}

	tests := []testCase[ TODO: Insert concrete types here ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AddAny(tt.x, tt.y); got != tt.want {
				t.Errorf("AddAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/

func TestAddAny(t *testing.T) {

	type testCase[T Addable] struct {
		name string
		x    T
		y    T
		want T
	}

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
