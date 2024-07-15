package utils

func Dedup[T comparable](data []T) []T {
	hashset := make(map[T]struct{})

	for _, d := range data {
		hashset[d] = struct{}{}
	}

	out := make([]T, 0)
	for k := range hashset {
		out = append(out, k)
	}

	return out
}
