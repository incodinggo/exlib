package sort_test

import (
	"fmt"
	"github.com/incodinggo/exlib/sort"
	"testing"
)

func TestSort(t *testing.T) {
	var arr1 = []int{9, 2, 4, 3, 5, 7, 6, 8, 1}
	var arr2 = []int64{9, 2, 4, 3, 5, 7, 6, 8, 1}
	var arr3 = []int64{9, 2, 4, 3, 5, 7, 6, 8, 1}
	var arr4 = []int64{9, 2, 4, 3, 5, 7, 6, 8, 1}
	var arr5 = []int64{9, 2, 4, 3, 5, 7, 6, 8, 1}
	var quickSort = sort.New[int](sort.Quick)
	var bubbleSort = sort.New[int64](sort.Bubble)
	var shellSort = sort.New[int64](sort.Shell)
	var mergeSort = sort.New[int64](sort.Merge)
	var heapSort = sort.New[int64](sort.Heap)
	quickSort.Sort(arr1)
	fmt.Println(arr1)
	bubbleSort.Sort(arr2)
	fmt.Println(arr2)
	shellSort.Sort(arr3)
	fmt.Println(arr3)
	mergeSort.Sort(arr4)
	fmt.Println(arr4)
	heapSort.Sort(arr5)
	fmt.Println(arr5)

}

func TestReverse(t *testing.T) {
	var arr1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sort.Reverse[int](arr1)
	fmt.Println(arr1)
}
