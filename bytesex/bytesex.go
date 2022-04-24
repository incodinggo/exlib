package bytesex

import (
	"fmt"
	"strings"
)

//byte
func String(b []byte) string {
	if len(b) == 0 {
		return ""
	}
	var s strings.Builder
	for _, a := range b {
		s.WriteString(fmt.Sprint(a))
	}
	return s.String()
}

func Strings(b []byte) []string {
	var s []string
	if len(b) == 0 {
		return s
	}
	for _, a := range b {
		s = append(s, fmt.Sprint(a))
	}
	return s
}
