package slice_test

import (
	"exlib/slice"
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	o := []int{1, 2, 3, 4, 5}
	slice.Insert(&o, []int{99}, 3, reflect.Int)
	fmt.Println(o)
	f := reflect.Swapper(o)
	f(2, 3)
}
