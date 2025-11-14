package piscine

func Abort(a, b, c, d, e int) int {
	values := []int{a, b, c, d, e}
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values)-1-i; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
			}
		}
	}
	return values[len(values)/2]
}
