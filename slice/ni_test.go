package slice_test

import (
	"fmt"
	"github.com/incodinggo/exlib/slice"
	"testing"
)

func TestNi(t *testing.T) {
	a := []string{"1", "3", "4", "5"}
	b := []string{"1", "3"}
	fmt.Println(slice.NI(a, b))
}
