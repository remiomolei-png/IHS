
import (
	"strconv"
)

func FromTo(from, to int) string {
	if (from > 99 || from < 0) || (to > 99 || to < 0) {
		return "Invalid"
	}
	res := ""
	if from <= to {
		for i := from; i <= to; i++ {
			if i != from {
				res += ", "
			}
			if i == 10 {
				res += strconv.Itoa(i)
			} else {
				res += "0" + strconv.Itoa(i)
			}
		}
	} else {
		for i := from; i >= to; i-- {
			if i != from {
				res += ", "
			}
			if i == 10 {
				res += strconv.Itoa(i)
			} else {
				res += "0" + strconv.Itoa(i)
			}
		}
	}
	return res + "\n"
}
