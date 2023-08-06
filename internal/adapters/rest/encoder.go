package rest

import (
	"math"
	"strings"
)

const (
	characterSet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	base         = 62
)

// ToBase62 encodes the given number to a base62 string.
func ToBase62(id int) string {
	r := id % base
	result := string(characterSet[r])
	div := id / base
	q := int(math.Floor(float64(div)))
	for q != 0 {
		r = q % base
		temp := q / base
		q = int(math.Floor(float64(temp)))
		result = string(characterSet[int(r)]) + result
	}
	return string(result)
}

// ToBase10 decodes a given base62 string to number to be used as a database ID.
func ToBase10(encodedString string) int {
	result := 0
	for _, r := range encodedString {
		result = (base * result) + strings.Index(characterSet, string(r))
	}
	return result
}
