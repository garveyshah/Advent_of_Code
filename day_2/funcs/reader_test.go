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
		{"Test 1", "../test_data/data1.txt", []string{"4x23x21", "22x29x19", "11x4x11", "8x10x5", "24x18x16", "11x25x22", "2x13x20", "24x15x14", "14x22x2", "30x7x3"}, nil},
		{"Test 2", "../test_data/data2.txt", []string{"12x9x18", "7x17x14", "13x17x17", "3x21x10", "30x9x15", "2x8x15", "15x12x10", "23x26x9", "29x30x10", "30x22x17", "17x26x30", "27x26x20", "17x28x17", "30x12x16", "7x23x15", "30x15x19", "13x19x10"}, nil},
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
		{"Test 1", []string{"4x23x21", "22x29x19", "11x4x11", "8x10x5", "24x18x16", "11x25x22", "2x13x20", "24x15x14", "14x22x2", "30x7x3"}, [][]int{{4, 23, 21}, {22, 29, 19}, {11, 4, 11}, {8, 10, 5}, {24, 18, 16}, {11, 25, 22}, {2, 13, 20}, {24, 15, 14}, {14, 22, 2}, {30, 7, 3}}},
		{"Test 2", []string{"12x9x18", "7x17x14", "13x17x17", "3x21x10", "30x9x15", "2x8x15", "15x12x10", "23x26x9", "29x30x10", "30x22x17", "17x26x30", "27x26x20", "17x28x17", "30x12x16", "7x23x15", "30x15x19", "13x19x10"}, [][]int{{12, 9, 18}, {7, 17, 14}, {13, 17, 17}, {3, 21, 10}, {30, 9, 15}, {2, 8, 15}, {15, 12, 10}, {23, 26, 9}, {29, 30, 10}, {30, 22, 17}, {17, 26, 30}, {27, 26, 20}, {17, 28, 17}, {30, 12, 16}, {7, 23, 15}, {30, 15, 19}, {13, 19, 10}}},
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
