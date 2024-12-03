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
		{"Test 1", "../test_data/data.txt", []string{"sszojmmrrkwuftyv", "isaljhemltsdzlum", "fujcyucsrxgatisb", "qiqqlmcgnhzparyg", "oijbmduquhfactbc", "jqzuvtggpdqcekgk", "zwqadogmpjmmxijf", "uilzxjythsqhwndh", "gtssqejjknzkkpvw", "wrggegukhhatygfi", "vhtcgqzerxonhsye", "tedlwzdjfppbmtdx", "iuvrelxiapllaxbg", "feybgiimfthtplui", "qxmmcnirvkzfrjwd", "vfarmltinsriqxpu"}, nil},
		{"Test 2", "../test_data/data1.txt", []string{"oanqfyqirkraesfq", "xilodxfuxphuiiii", "yukhnchvjkfwcbiq", "bdaibcbzeuxqplop", "ivegnnpbiyxqsion", "ybahkbzpditgwdgt", "dmebdomwabxgtctu", "ibtvimgfaeonknoh", "jsqraroxudetmfyw", "dqdbcwtpintfcvuz", "tiyphjunlxddenpj"}, nil},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Reader(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("%q Failed: got %q, want %q", tc.name, got, tc.want)
			}
			if err != tc.err {
				t.Errorf("%q Failed: got %v, want %v", tc.name, err, tc.err)
			}
		})
	}
}
