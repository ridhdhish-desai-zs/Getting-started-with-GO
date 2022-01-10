package testing

import "strings"

func isPalindrom(str string) bool {
	str = strings.ToLower(str)

	i, j := 0, len(str)-1

	for i < len(str)/2 {

		if str[i] < 97 || str[i] > 122 {
			i++
			continue
		}

		if str[j] < 97 || str[j] > 122 {
			j--
			continue
		}

		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}
