package sort

// SelectionSort 选择排序
func SelectionSort(x []int) {
	if len(x) <= 1 {
		return
	}

	for i := 0; i < len(x); i++ {
		j := i + 1
		minIndex := i
		for j < len(x) {

			if x[j] < x[minIndex] {
				minIndex = j
			}
			j++
		}
		x[i], x[minIndex] = x[minIndex], x[i]
	}
}
