package slice_test

import (
	"fmt"
	"github.com/incodinggo/exlib/slice"
	"testing"
)

func TestInsert(t *testing.T) {
	o := []int{1, 2, 3, 4, 5}
	slice.Insert[int](&o, []int{99}, 3)
	fmt.Println(o)
	o1 := []int{1, 2, 3, 4, 5, 0, 5, 0, 2}
	fmt.Println(slice.RmDup(o1, true))
}
