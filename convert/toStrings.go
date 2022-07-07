package convert

import (
	"fmt"
)

func Strings[T any](in []T) []string {
	var out []string
	for _, e := range in {
		out = append(out, fmt.Sprint(e))
	}
	return out
}
