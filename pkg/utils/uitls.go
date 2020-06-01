package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func FormatByteSlice(s []byte) string {
	representation := make([]string, len(s))
	for i, v := range s {
		representation[i] = fmt.Sprintf("0x%x", v)
	}
	return strings.Join(representation, ",")
}

func ParseByteSlice(r string) []byte {
	elements := strings.Split(strings.ReplaceAll(r, " ", ""), ",")

	s := make([]byte, len(elements))
	for i, v := range elements {
		d, _ := strconv.ParseUint(v, 16, 64)
		s[i] = byte(d)
	}

	return s
}
