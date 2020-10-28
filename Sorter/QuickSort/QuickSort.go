package quicksort

func quickSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	if len(values) == 2 {
		if values[0] > values[1] {
			values[0], values[1] = values[1], values[0]
		}
		return values
	}
	left := 0
	camp := len(values) - 1
	right := camp - 1
	for left < right {
		for values[left] <= values[camp] && left < right {
			left++
		}
		for values[right] >= values[camp] && right > left {
			right--
		}
		if left < right {
			values[left], values[right] = values[right], values[left]
		}
		if values[camp] < values[right] {
			values[right], values[camp] = values[camp], values[right]
			camp = right
		}
	}
	quickSort(values[:camp])
	quickSort(values[camp:])
	return values
}

func QuickSort(values []int) []int {
	return quickSort(values)
}

/*上面是采用交换数字，最后移动轴
也可以直接移动轴
*/
func quickSortNew1(values []int) {
	if len(values) <= 1 {
		return
	}
	if len(values) == 2 {
		if values[0] > values[1] {
			values[0], values[1] = values[1], values[0]
		}
		return
	}
	var pointer, left, right int
	pointer = len(values) - 1
	left = 0
	right = pointer
	for left < right {
		for left < pointer && values[left] <= values[pointer] {
			left++
		}
		if left <= pointer {
			values[left], values[pointer] = values[pointer], values[left]
			pointer = left
		}
		for right > pointer && values[right] >= values[pointer] {
			right--
		}
		if right >= pointer {
			values[right], values[pointer] = values[pointer], values[right]
			pointer = right
		}
	}
	quickSortNew1(values[:pointer])
	quickSortNew1(values[pointer:])
}

func QuickSortNew1(values []int) {
	quickSortNew1(values)
}
