package sort

type bubbleSort[T Type] struct {
}

func (s *bubbleSort[T]) Sort(sl []T) {
	s.sort(sl)
}

func (s *bubbleSort[T]) sort(sl []T) {
	l := len(sl)
	for i := 0; i < l; i++ {
		for j := 0; j < l-i-1; j++ {
			if sl[j] > sl[j+1] {
				sl[j+1], sl[j] = sl[j], sl[j+1]
			}
		}
	}
}
