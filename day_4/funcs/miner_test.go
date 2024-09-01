package funcs

import "testing"

func TestMD5(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  string
	}{
		// {"Test 1", "abcdef", 609043},
		// {"Test 2", "pqrstuv", 1048970},
		//	{"Test 3","bgvyzdsv",1048970},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := MD5(tc.input)
			if got != tc.want {
				t.Errorf("%q Failed: got %v, want %v", tc.name, got, tc.want)
			}
		})
	}
}
