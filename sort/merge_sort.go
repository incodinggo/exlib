package sort

type mergeSort[T Type] struct {
}

func (s *mergeSort[T]) Sort(sl []T) {
	s.sort(sl, 0, len(sl)-1)
}

func (s *mergeSort[T]) sort(sl []T, l, r int) {
	if l >= r {
		return
	}
	mid := (r + l) / 2
	s.sort(sl, l, mid)
	s.sort(sl, mid+1, r)

	merge(sl, l, mid, r)
}

func merge[T Type](sl []T, l, mid, r int) {
	tmp := make([]T, r-l+1)
	for i := l; i <= r; i++ {
		tmp[i-l] = sl[i]
	}

	left := l
	right := mid + 1

	for i := l; i <= r; i++ {
		if left > mid {
			sl[i] = tmp[right-l]
			right++
		} else if right > r {
			sl[i] = tmp[left-l]
		} else if tmp[left-l] > tmp[right-l] {
			sl[i] = tmp[right-l]
			right++
		} else {
			sl[i] = tmp[left-l]
			left++
		}
	}
}
