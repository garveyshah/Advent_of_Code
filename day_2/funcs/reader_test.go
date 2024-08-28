package funcs

import (
	"reflect"
	"testing"
)

func TestReader(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  []string
		err   error
	}{
		// {"Test 1", "/test_data/data1.txt", []string{}, nil},
		// {"Test 2", "/test_data/data2.txt", []string{}, nil},
		// {"Test 3", "/test_data/data3.txt", []string{}, nil},
		// {"Test 4", "/test_data/data4.txt", []string{}, nil},
		// {"Test 5", "/test_data/data5.txt", []string{}, nil},
		// {"Test 6", "/test_data/data6.txt", []string{}, nil},
		// {"Test 7", "/test_data/data7.txt", []string{}, nil},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Reader(tc.input)
			if err != tc.err {
				t.Errorf("%q, Failed : got %q want %q", tc.name, err, tc.err)
			}
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%q, Failed: got %q want %q", tc.name, got, tc.want)
			}
		})
	}
}

func TestDecodeData(t *testing.T) {
	tt := []struct {
		name  string
		input []string
		want  [][]int
	}{
		// {"Test 1", []string{}, [][]int{}},
		// {"Test 2", []string{}, [][]int{}},
		// {"Test 3", []string{}, [][]int{}},
		// {"Test 4", []string{}, [][]int{}},
		// {"Test 5", []string{}, [][]int{}},
		// {"Test 6", []string{}, [][]int{}},
		// {"Test 7", []string{}, [][]int{}},
		// {"Test 8", []string{}, [][]int{}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := DecodeData(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%q, Failed: got %q want %q", tc.name, got, tc.want)
			}
		})
	}
}
