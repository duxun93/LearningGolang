package countsort

//对于已知范围的整数数组，且变化范围不大（且已知）的同时数组最小值从0开始，用直接计数
//不稳定
func countSortByCount(values []int, rangeValues int) []int {
	var count, result []int
	for i := 0; i < rangeValues+1; i++ {
		count = append(count, 0)
	}
	for index, num := range values {
		count[num]++
	}
	for num, counts := range count {
		for i := 0; i < counts; i++ {
			result = append(result, num)
		}
	}
	return result

}
func CountSortByCount(values []int, rangeValues int) {
	return countSortByCount(values, rangeValues)
}

//数组范围已知，且不从0开始
func countSortBySubtraction(values []int, rangeValues []int) []int {

}
