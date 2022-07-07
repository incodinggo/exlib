package sort

type shellSort[T Type] struct {
}

func (s *shellSort[T]) Sort(sl []T) {
	s.sort(sl)
}

func (s *shellSort[T]) sort(sl []T) {
	l := len(sl)
	h := 1
	for h < l/3 {
		h = 3*h + 1
	}

	for h >= 1 {
		for i := h; i < l; i++ {
			for j := i; j >= h && sl[j] < sl[j-1]; j -= h {
				sl[j], sl[j-1] = sl[j-1], sl[j]
			}
		}
		h /= 3
	}
}
