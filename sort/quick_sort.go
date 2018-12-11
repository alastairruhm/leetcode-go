package sort

// QuickSort 快排
func QuickSort(a []int) {
	// temp := make([]int, len(a))
	quicksort(a, 0, len(a)-1)
}

func quicksort(arr []int, left int, right int) {
	if left < right {
		pivot := partition(arr, left, right)
		quicksort(arr, left, pivot-1)
		quicksort(arr, pivot+1, right)
	}
}

func partition(arr []int, left int, right int) int {
	pivot := arr[right]
	i := left
	for j := left; j < right; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[right] = arr[right], arr[i]
	// fmt.Printf("%+v\n", arr)
	return i
}
