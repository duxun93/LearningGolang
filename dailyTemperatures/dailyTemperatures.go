/*
leetcode 739 Daily Temperature
Day 20-12-06
*/

package dailtTemperature

func dailyTemperatures(T []int) []int {
	var result []int
	result = []int{}
	for index, _ := range T {
		result = append(result, checkTempure(&T, index))
	}
	return result
}

func checkTempure(node *[]int, n int) (day int) {
	temperature := (*node)[n]
	for index := n + 1; index < len(*node); index++ {
		if (*node)[index] > temperature {
			day = index - n
			return day
		}
	}
	return 0
}
