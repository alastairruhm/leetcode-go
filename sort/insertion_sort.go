package sort

import (
	"fmt"
)

// InsertionSort 插入排序
func InsertionSort(x []int) {
	if len(x) <= 1 {
		return
	}

	for i := 1; i < len(x); i++ {
		value := x[i] // 待插入元素
		j := i - 1
		for j >= 0 {
			if x[j] > value {
				x[j+1] = x[j]
			} else {
				break
			}
			j--
		}
		fmt.Println(j)
		x[j+1] = value
	}
}
