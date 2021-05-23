package numbers

func Even(n int) bool {
	if n == 0 {
		return true
	}
	return n%2 == 0
}
