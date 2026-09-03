package even

func IsEven(a int) bool {
	if a < 0 {
		a = a * (-1)
	}
	if a%2 == 1 {
		return false
	} else {
		return true
	}
}
