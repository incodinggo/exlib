package sort

type quickSort[T Type] struct {
}

func (s *quickSort[T]) Sort(sl []T) {
	s.sort(sl, 0, len(sl)-1)
}

func (s *quickSort[T]) sort(sl []T, left, right int) {
	if left < right {
		k := sl[(left+right)/2]
		i := left
		j := right
		for {
			for sl[i] < k {
				i++
			}
			for sl[j] > k {
				j--
			}
			if i >= j {
				break
			}
			sl[i], sl[j] = sl[j], sl[i]
		}
		s.sort(sl, left, i-1)
		s.sort(sl, j+1, right)
	}
}
