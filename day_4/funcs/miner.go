package funcs

func MD5(s string) string {
	return "Godwin"
}

// StringBuilder to concatenate secret key and number
func StringBuilder(sKey string, num int) string {
	return sKey + CustomAtoi(num)
}

// CustomAtoi to convert interger to string
func CustomAtoi(num int) string {
	var result string

	for num > 0 {
		result = string('0'+(num%10)) + result
		num /= 10
	}

	if result == "" {
		return "0"
	}

	return result
}
