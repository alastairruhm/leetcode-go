package reverseInteger

import "math"

// MIN min int
var MIN = 0x80000000

// MAX max int
var MAX = 0x7FFFFFFF

func reverse(x int) int {
	sum := 0
	for {
		leftDigits := x / 10
		lastDigit := x % 10
		x = leftDigits

		sum = sum*10 + lastDigit

		if 0 == leftDigits {
			break
		}
	}

	if sum < -MIN || sum > MAX {
		sum = 0
	}

	return sum
}

func reverseIn3msSolution(x int) int {
	sum := 0
	for {
		leftDigits := x / 10
		lastDigit := x % 10
		x = leftDigits

		sum = sum*10 + lastDigit

		if 0 == leftDigits {
			break
		}
	}

	if sum < -MIN || sum > MAX {
		sum = 0
	}

	return sum
}

func reverseWithGoroutine(x int) int {
	ch := make(chan int)
	go Parse(x, ch)
	var ns []int
	for n := range ch {
		ns = append(ns, n)
	}
	size := len(ns)
	sum := 0
	for index, digit := range ns {
		d := int(math.Pow10(size - index - 1))
		sum = sum + digit*d
	}

	if sum < -MIN || sum > MAX {
		sum = 0
	}

	return sum
}

// ParseRecurse parse int with every number from low to high recursivly
func ParseRecurse(x int, ch chan int) {
	remain := x / 10
	n := x % 10
	ch <- n
	if remain != 0 {
		ParseRecurse(remain, ch)
	}
}

// Parse ...
func Parse(x int, ch chan int) {
	ParseRecurse(x, ch)
	close(ch)
}
