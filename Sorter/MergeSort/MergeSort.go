package mergesort

//递归方法归并排序
func mergeSort(values []int) []int {
	if len(values) <= 1 {
		return values
	}
	if len(values) == 2 {
		if values[0] > values[1] {
			values[0], values[1] = values[1], values[0]
		}
		return values
	}
	var left, right, result []int
	left = mergeSort(values[:len(values)>>1])
	right = mergeSort(values[len(values)>>1:])
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] <= right[rightIndex] {
			result = append(result, left[leftIndex])
			leftIndex++
		} else {
			result = append(result, right[rightIndex])
			rightIndex++
		}
	}
	if rightIndex < len(right) {
		result = append(result, right[rightIndex:]...)
	} else if leftIndex < len(right) {
		result = append(result, left[leftIndex:]...)
	}
	return result
}
func MergeSort(values []int) []int {
	return mergeSort(values)
}

//循环方法归并排序
