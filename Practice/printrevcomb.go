func PrintRevComb() {
	result := ""
	first := true
	
	// Start from 9 down to 2 (first digit)
	for i := 9; i >= 2; i-- {
		// Second digit must be less than first
		for j := i - 1; j >= 1; j-- {
			// Third digit must be less than second
			for k := j - 1; k >= 0; k-- {
				if !first {
					result += ", "
				}
				first = false
				result += fmt.Sprintf("%d%d%d", i, j, k)
			}
		}
	}
	
	fmt.Println(result)
}
