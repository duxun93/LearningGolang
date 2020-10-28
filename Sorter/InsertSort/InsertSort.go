package insertsort

func insertSort(values []int) {
	for index := 1; index < len(values); index++ {
		for num := 0; num < index; num++ {
			if values[num] > values[index] {
				values[num], values[index] = values[index], values[num]
			}
		}
	}
}

func InsertSort(values []int) {
	insertSort(values)
}
