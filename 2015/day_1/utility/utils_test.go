package utility

import "testing"

func TestGetData(t *testing.T) {
	Tests := []struct {
		name string
		data string
		err  error
	}{
		{"Test1", "(((()))()(()()()(", nil},
	}

	for _, tt := range Tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetData("data.txt")
			if got != tt.data {
				t.Errorf("%#v got %s want %s", tt.name, got, tt.data)
			}
		})
	}
}
