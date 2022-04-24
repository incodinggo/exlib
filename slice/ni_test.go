package slice_test

import (
	"ExLib/slice"
	"fmt"
	"testing"
)

func TestNi(t *testing.T)  {
	a := []string{"1", "3", "4", "5"}
	b := []string{"1", "3"}
	fmt.Println(slice.NIStr(a, b).Strings())
}
