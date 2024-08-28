package funcs

import "testing"

func TestWrappPaper(t *testing.T) {
	tt := []struct {
		name  string
		input []int
		want  int
	}{
		{"Test 1", []int{2, 3, 4}, 58},
		{"Test 2", []int{1, 1, 10}, 43},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := WrappPaper(tc.input)
			if got != tc.want {
				t.Errorf("%q Failed: got %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}

func TestTotalSa(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"Test 1", [][]int{{2, 3, 4}, {1, 1, 10}}, 101},
		{"Test 2", [][]int{{2, 3, 4}, {1, 1, 10}}, 101},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := TotalSa(tc.input)
			if got != tc.want {
				t.Errorf("%q Failed: got %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	tt := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{"Test 1", 2, 3, 4, 2},
		{"Test 2", 1, 1, 10, 1},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Min(tc.a, tc.b, tc.c)
			if got != tc.want {
				t.Errorf("%q Failed: got %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}

func TestRibbon(t *testing.T) {
	tt := []struct {
		name  string
		input [][]int
		want  int
	}{
		{"Test 1", [][]int{{2, 3, 4}, {1, 1, 10}}, 48},
		{"Test 2", [][]int{{2, 3, 4}, {1, 1, 10}}, 48},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := Ribbon(tc.input)
			if got != tc.want {
				t.Errorf("%q Failed: got %q, want %q", tc.name, got, tc.want)
			}
		})
	}
}
