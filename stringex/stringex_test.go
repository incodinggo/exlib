package stringex_test

import (
	"fmt"
	"github.com/incodinggo/exlib/stringex"
	"testing"
)

func TestStringex(t *testing.T) {
	fmt.Println(stringex.Join([]int{1, 2, 3, 4, 5, 6}, ";"))
}
