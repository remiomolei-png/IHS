package main

func IsPrime(nb int) {
	if nb < 2{
		return false
	}
	for a:=2; a*a<=nb; a++ {
		if nb%a == 0{
			return false
		}
	}
	return true
}

func SumPrimes(n int) {
	sum := 0
	for i := 2; i <=n; i++ {
		if IsPrime(i){
			sum += i
		}
	}
	return sum
} 
