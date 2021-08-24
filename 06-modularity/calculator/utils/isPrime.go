package utils

func IsPrime(no int) bool {
	if no < 2 {
		return false
	}
	for i := 2; i <= no/2; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
