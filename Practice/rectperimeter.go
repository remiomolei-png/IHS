package piscine

func RectPerimeter(w, h int) int {
	if w <=0 || h <= 0 {
		return -1
	}
	perimeter := (w + h)*2
	return perimeter
}
