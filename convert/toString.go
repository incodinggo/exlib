package convert

import (
	"fmt"
)

func I2S(in []int) []string {
	var out []string
	for _, e := range in {
		out = append(out, fmt.Sprint(e))
	}
	return out
}

func I2S64(in []int64) []string {
	var out []string
	for _, e := range in {
		out = append(out, fmt.Sprint(e))
	}
	return out
}

func Inf2S(in []interface{}) []string {
	var out []string
	for _, e := range in {
		out = append(out, fmt.Sprint(e))
	}
	return out
}
