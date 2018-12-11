package sort

// MergeSort 归并排序
func MergeSort(a []int) {
	if (len(a)) <= 1 {
		return
	}
	temp := make([]int, len(a))
	mergesort(a, 0, len(a)-1, temp)
}

func mergesort(arr []int, left int, right int, temp []int) {
	if left < right {
		mid := (left + right) / 2
		mergesort(arr, left, mid, temp)    // 左边归并排序，使得左子序列有序
		mergesort(arr, mid+1, right, temp) // 右边归并排序，使得右子序列有序
		merge(arr, left, mid, right, temp) // //将两个有序子数组合并操作
	}
}

func merge(arr []int, left int, mid int, right int, temp []int) {
	// fmt.Println(left, mid, right)
	i := left    // left array pointer
	j := mid + 1 // right array pointer
	t := 0       // temp pointer
	for i <= mid && j <= right {
		if arr[i] <= arr[j] {
			temp[t] = arr[i]
			t++
			i++
		} else {
			temp[t] = arr[j]
			t++
			j++
		}
	}
	for i <= mid {
		temp[t] = arr[i]
		t++
		i++
	}
	for j <= right {
		temp[t] = arr[j]
		t++
		j++
	}

	t = 0
	for left <= right {
		arr[left] = temp[t]
		left++
		t++
	}
	// fmt.Printf("%+v\n", temp)
	// fmt.Printf("%+v\n", arr)
}
