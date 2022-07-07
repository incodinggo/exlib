package ternary_test

import (
	"fmt"
	"github.com/incodinggo/exlib/ternary"
	"testing"
)

func TestIf(t *testing.T) {
	fmt.Println(ternary.If(true, 1, 2))
}
