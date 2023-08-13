package rest

import (
	"strings"
)

const (
	characterSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base         = 62
)

// ToBase62 encodes the given number to a base 62 string.
func ToBase62(id int) string {
	remainder := id % base
	result := string(characterSet[remainder])
	div := id / base
	q := int(float64(div))
	for q != 0 {
		remainder = q % base
		temp := q / base
		q = int((float64(temp)))
		result = string(characterSet[int(remainder)]) + result
	}
	return string(result)
}

// ToBase10 decodes a given base 62 string to number to be used as a database ID.
func ToBase10(encodedString string) int {
	result := 0
	for _, r := range encodedString {
		result = (base * result) + strings.Index(characterSet, string(r))
	}
	return result
}
