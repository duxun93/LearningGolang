package shellsort

import "fmt"

func shellSort(values []int) {
	gap := getFnuthList(len(values))
	//fmt.Print(gap)
	for ; gap > 0; gap = (gap - 1) / 3 {
		fmt.Print(gap)
		for index := gap - 1; index < len(values); index++ {
			for num := 0; num < index; num = num + gap {
				if values[num] > values[index] {
					values[num], values[index] = values[index], values[num]
				}
				fmt.Print(values)
			}
		}

	}
}

func getFnuthList(lenth int) int {
	var i int
	for i = 1; 3*i < lenth-1; i = 3*i + 1 {
		continue
	}
	return i
}

func ShellSort(values []int) {
	shellSort((values))
}
