package testing

func reverseString(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func isPalindrom(str string) bool {
	revStr := reverseString(str)

	return revStr == str
}
