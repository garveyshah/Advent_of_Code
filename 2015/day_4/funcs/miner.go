package funcs

import (
	"crypto/md5"
	"strings"
)

// StringBuilder to concatenate secret key and number
func StringBuilder(sKey string, num int) string {
	return sKey + CustomAtoi(num)
}

// CustomAtoi to convert interger to string
func CustomAtoi(num int) string {
	if num == 0 {
		return "0"
	}
	var result string
	for num > 0 {
		r := rune(('0') + (num % 10))
		result = string(r) + result
		num /= 10
	}

	return result
}

// Md5Hash() generates the MD5 hash for a given message using a custom hex encoder.
func Md5Hash(sKey string) string {
	hash := md5.Sum([]byte(sKey))

	return ByteToHex(hash[:])
}

// HexToString() to convert a slcie to a hexadecimal string.
func ByteToHex(data []byte) string {
	hexChars := "0123456789abcdef"
	result := make([]byte, len(data)*2)
	for i, b := range data {
		result[i*2] = hexChars[b>>4]
		result[i*2+1] = hexChars[b&0x0F]
	}
	return string(result)
}

// AdvCoinCointMIner() finds the lowest number that produces a hash with five leading zeroes
func AdvCoinMiner(sKey string) int {
	for num := 1; ; num++ {
		hash := Md5Hash(StringBuilder(sKey, num))
		if strings.HasPrefix(hash, "000000") {
			return num
		}
	}
}
