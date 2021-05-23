package numbers

import "log"

func Even(n int) bool {
	if n == 0 {
		log.Println("handling zero value input")
		return true
	}
	return n%2 == 0
}
