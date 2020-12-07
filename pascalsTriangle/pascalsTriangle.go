/*
leetcode 118 杨辉三角
Day 20-12-06
*/
package pascalsTriangle

func generate(numRows int) [][]int {
	var result [][]int
	if numRows == 0 {
		return nil
	}
	result = [][]int{{1}, {1, 1}}
	if numRows == 1 {
		return [][]int{{1}}
	}
	if numRows == 2 {
		return [][]int{{1}, {1, 1}}
	}
	for i := 1; i < numRows-1; i++ {
		last := result[i]
		temp := []int{}
		temp = append(temp, 1)
		for index := 0; index < len(last)-1; index++ {
			temp = append(temp, last[index]+last[index+1])
		}
		temp = append(temp, 1)
		result = append(result, temp)
	}
	return result
}
