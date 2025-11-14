func FindPrevPrime(nb int) int {
	if nb < 2 {
		return 0
	}
	if is.Prime(nb) {
		return nb
	}
	return FindPrevPrime(nb - 1)
}

func IsPrime(nb int) bool {
    if nb <= 1 {
        return false
    }
    if nb == 2 {
        return true
    }
    if nb%2 == 0 {
        return false  // Quick check for even numbers
    }

    for i := 3; i*i <= nb; i += 2 {  // Skip even numbers
        if nb%i == 0 {
            return false
        }
    }
    return true
}
