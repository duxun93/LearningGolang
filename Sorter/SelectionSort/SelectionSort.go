//selection
package selectionsort

func selection(values []int) {
	for index := 0; index < len(values); index++ {
		var min int
		min = index
		for i := index + 1; i < len(values); i++ {
			if values[i] < values[min] {
				min = i
			}
		}
		values[index], values[min] = values[min], values[index]
	}
	return
}

func Selection(values []int) []int {
	selection(values)
	return values
}

/*给出了一个优化，每次遍历同时找最大和最小值;
  不使用额外数组空间 */
func selectionNew1(values []int) {
	//fmt.Print("start")
	for index := 0; index < len(values)/2; index++ {
		var min, max int
		min = index
		max = index
		for i := index + 1; i < len(values)-index; i++ {
			if values[i] < values[min] {
				min = i
			}
			if values[i] > values[max] {
				max = i
			}
		}
		values[index], values[min] = values[min], values[index]
		if max == index && min != len(values)-index-1 {
			values[len(values)-index-1], values[min] = values[min], values[len(values)-1-index]
		} else if max == index && min == len(values)-index-1 {
			continue
		} else {
			values[len(values)-index-1], values[max] = values[max], values[len(values)-1-index]
		}

	}
}

func SelectionNew1(values []int) []int {
	selectionNew1(values)
	return values
}

/*优化另一种方式，思路简单;
但是需要使用一个额外的数组*/
func selectionNew2(values []int) []int {
	var result []int
	result = make([]int, len(values))
	for index := 0; index < len(values)/2; index++ {
		var min, max int
		min = index
		max = index
		for i := index + 1; i < len(values)-index; i++ {
			if values[i] < values[min] {
				min = i
			}
			if values[i] > values[max] {
				max = i
			}
		}
		result[index] = values[min]
		result[len(values)-index-1] = values[max]
	}
	return result
}

func SelectionNew2(values []int) []int {
	return selectionNew2(values)
}
