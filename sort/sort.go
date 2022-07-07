package sort

import "golang.org/x/exp/constraints"

type Type interface {
	constraints.Ordered
}

type Inf[T Type] interface {
	Sort([]T)
}

const (
	Quick = iota
	Bubble
	Shell
	Merge
	Heap
)

func New[T Type](method int) Inf[T] {
	switch method {
	case Quick:
		return new(quickSort[T])
	case Bubble:
		return new(bubbleSort[T])
	case Shell:
		return new(shellSort[T])
	case Merge:
		return new(mergeSort[T])
	case Heap:
		return new(heapSort[T])
	default:
		return nil
	}
}

func Reverse[T Type](sl []T) {
	for i, j := 0, len(sl)-1; i < j; i, j = i+1, j-1 {
		sl[i], sl[j] = sl[j], sl[i]
	}
}
