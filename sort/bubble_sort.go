package sort

// BubbleSort 冒泡排序
func BubbleSort(x []int) {
	if len(x) <= 0 {
		return
	}

	for i := 0; i < len(x); i++ {

		flag := false

		for j := 0; j < len(x)-i-1; j++ {
			if x[j] > x[j+1] {
				x[j], x[j+1] = x[j+1], x[j] // swap
				flag = true
			}
		}

		if !flag {
			break
		}

	}
}
