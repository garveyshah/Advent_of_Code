package funcs

import "testing"

// Test for MD5()
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

// Test for StringBuilder()
func TestStringBuilder(t *testing.T) {
	tt := []struct {
		name     string
		inputStr string
		inputInt int
		want     string
	}{
		{"Test 1", "My age is ", 25, "My age is 25"},
		{"Test 2", "567", 8910, "5678910"},
		{"Test 3", "My costs  ", 25, "My car costs 25"},
		{"Test 4", "01234", 5678, "012345678"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := StringBuilder(tc.inputStr, tc.inputInt)
			if got != tc.want {
				t.Errorf("%q, Failed: got %v want %v", tc.name, got, tc.want)
			}
		})
	}
}
