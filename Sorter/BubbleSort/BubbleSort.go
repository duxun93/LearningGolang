//bubblesort
package bubblesort

import "fmt"

func bubbleSort(values []int) {
	var flag bool
	flag = true
	for flag {
		count := 0
		for index := 0; index < len(values)-1; index++ {
			if values[index] > values[index+1] {
				values[index], values[index+1] = values[index+1], values[index]
				count++
			}
		}
		if count == 0 {
			flag = false
		}
	}
}

func BubbleSort(values []int) {
	bubbleSort(values)
}

//优化，其实冒泡排序相当于选择排序的一种，选择排序找最小值往前排。
//冒泡排序是找最大的往后移

func bubbleSortNew(values []int) {
	for count := 0; count < len(values)-1; count++ {
		for index := 0; index < len(values)-count-1; index++ {
			if values[index] > values[index+1] {
				values[index], values[index+1] = values[index+1], values[index]
			}
		}
	}
}

func BubbleSortNew(values []int) {
	bubbleSortNew(values)
}
