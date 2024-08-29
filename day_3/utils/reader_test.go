package utils

import "testing"

func TestReader(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
		err   error
	}{
		{"Test 1", "../test_data/data1.txt", ">", nil},
		{"Test 2", "../test_data/data2.txt", "^>v<", nil},
		{"Test 3", "../test_data/data3.txt", "^v^v^v^v^v", nil},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Reader(tc.input)
			if err != tc.err {
				t.Errorf("%q, Failed logged wrong error : got %v, want %v", tc.name, got, tc.err)
			}
			if got != tc.want {
				t.Errorf("%q, Failed: got %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
