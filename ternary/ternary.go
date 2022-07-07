package ternary

func If[T any](ok bool, t, f T) T {
	if ok {
		return t
	}
	return f
}
