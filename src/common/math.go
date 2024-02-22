package common

func Clamp(a, b, c int) int {
	return min(c, max(a, b))
	// return min(max(a, b), c)
}
