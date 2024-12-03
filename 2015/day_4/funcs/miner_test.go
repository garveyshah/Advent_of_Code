package funcs

import "testing"

// Test for AdvCoinMiner()
func TestAdvCoinMiner(t *testing.T) {
	tt := []struct {
		name  string
		input string
		want  int
	}{
		// {"Test 1", "abcdef", 609043},
		// {"Test 2", "pqrstuv", 1048970},
		//	{"Test 3","bgvyzdsv",1048970},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := AdvCoinMiner(tc.input)
			if got != tc.want {
				t.Errorf("%q Failed: got %d, want %d", tc.name, got, tc.want)
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
		{"Test 3", "My car costs ", 25, "My car costs 25"},
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

// Test for CustomAtoi()
func TestCustomAtoi(t *testing.T) {
	tt := []struct {
		name     string
		inputInt int
		want     string
	}{
		{"Test 1", 25, "25"},
		{"Test 2", 567, "567"},
		{"Test 3", 5, "5"},
		{"Test 4", 5678, "5678"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got := CustomAtoi(tc.inputInt)
			if got != tc.want {
				t.Errorf("%q, Failed: got %v want %v", tc.name, got, tc.want)
			}
		})
	}
}

// Test for Md5Hash
func TestMd5Hash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Test 1", "hello", "5d41402abc4b2a76b9719d911017c592"}, // MD5 hash of "hello"
		{"Test 2", "world", "7d793037a0760186574b0282f2f435e7"}, // MD5 hash of "world"
		//{"Test 3", "bgvyzdsv1", "098f6bcd4621d373cade4e832627b4f6"}, // Test case with specific key
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := Md5Hash(test.input)
			if output != test.expected {
				t.Errorf("%q Failed : Md5Hash(%s) = %s; expected %s", test.name, test.input, output, test.expected)
			}

		})
	}
}

// Test for ByteToHex
func TestByteToHex(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected string
	}{
		{"Test 1", []byte{0x5d, 0x41, 0x40, 0x2a, 0xbc, 0x4b, 0x2a, 0x76, 0xb9, 0x71, 0x9d, 0x91, 0x10, 0x17, 0xc5, 0x92}, "5d41402abc4b2a76b9719d911017c592"},
		{"Test 2", []byte{0x7d, 0x79, 0x30, 0x37, 0xa0, 0x76, 0x01, 0x86, 0x57, 0x4b, 0x02, 0x82, 0xf2, 0xf4, 0x35, 0xe7}, "7d793037a0760186574b0282f2f435e7"},
		{"Test 3", []byte{0x09, 0x8f, 0x6b, 0xcd, 0x46, 0x21, 0xd3, 0x73, 0xca, 0xde, 0x4e, 0x83, 0x26, 0x27, 0xb4, 0xf6}, "098f6bcd4621d373cade4e832627b4f6"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			output := ByteToHex(test.input)
			if output != test.expected {
				t.Errorf("%q Failed :- ByteToHex(%v) = %s; expected %s", test.name, test.input, output, test.expected)
			}
		})
	}
}
