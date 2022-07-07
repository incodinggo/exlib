package sort

type heapSort[T Type] struct {
}

func (s *heapSort[T]) Sort(sl []T) {
	s.sort(sl)
}

func (s *heapSort[T]) sort(sl []T) {
	l := len(sl)
	for i := l / 2; i > -1; i-- {
		heap(sl, i, l-1)
	}
	for i := l - 1; i > 0; i-- {
		sl[i], sl[0] = sl[0], sl[i]
		heap(sl, 0, i-1)
	}
}

func heap[T Type](sl []T, i, end int) {
	l := 2*i + 1
	if l > end {
		return
	}
	n := l
	r := 2*i + 2
	if r <= end && sl[r] > sl[l] {
		n = r
	}
	if sl[i] > sl[n] {
		return
	}
	sl[n], sl[i] = sl[i], sl[n]
	heap(sl, n, end)
}
