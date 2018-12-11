// 如何编程实现“求一个数的平方根”？要求精确到小数点后 6 位。

package c001

import (
	"math"
)

func Sqrt(n float64, precision float64) float64 {
	if n < 0 {
		panic("X")
	}

	low := float64(0)
	up := float64(n)
	mid := float64((low + up) / 2)
	for math.Abs(mid*mid-n) > precision {
		if mid*mid > n {
			up = mid
			mid = low + float64((mid-low)/2)
		} else {
			low = mid
			mid = up - (up-mid)/2
		}
	}

	return mid
}
