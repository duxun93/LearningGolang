package radixsort

import (
	"math"
)

func getNum(num int, position int) (result int) {
	result = (num / int(math.Pow10(position))) % int(math.Pow10(position+1))
	return result
}

func bottleSort(values []int, position int) (result []int, flag bool) {
	count := [10][]int{}
	for i := 0; i < len(values); i++ {
		num := getNum(values[i], position)
		count[num] = append(count[num], values[i])
	}
	if len(count[0]) == len(values) {
		flag = true
		result = values
		return
	} else {
		flag = false
		for i := 0; i < len(count); i++ {
			if len(count[i]) > 0 {
				result = append(result, count[i]...)
			}
		}
	}
	return

}

func radixSort(values []int) []int {
	result := []int{}
	result = append(result, values...)
	flag := false
	//count := make([][]int)
	for i := 0; ; i++ {
		if flag {
			return result
		} else {
			result, flag = bottleSort(result, i)
		}
	}
}

func RadixSort(values []int) []int {
	result := radixSort(values)
	return result
}

//基数排序考虑字符串排序
