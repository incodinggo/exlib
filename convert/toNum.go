package convert

import "strconv"

func Int64(in []string) []int64 {
	var out []int64
	for _, e := range in {
		i, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil
		}
		out = append(out, i)
	}
	return out
}

func Int(in []string) []int {
	var out []int
	for _, e := range in {
		i, err := strconv.ParseInt(e, 10, 64)
		if err != nil {
			return nil
		}
		out = append(out, int(i))
	}
	return out
}
