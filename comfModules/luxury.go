package comfModules


func IsPrime(number int) bool {
	if number <= 1 {
		return false
	}
	for i := 2; i*i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}


func Swap(a int, b int) {
	tmp := a
	a = b
	b = tmp
}
