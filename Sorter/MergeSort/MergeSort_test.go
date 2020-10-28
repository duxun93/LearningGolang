//mergeSort test
//个人觉得测试这种排序最好的办法其实是用质数个数测试。
package mergesort

import "testing"

func TestInsertSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	result := MergeSort(values)
	if result[0] != 1 || result[1] != 2 || result[2] != 3 || result[3] != 4 || result[4] != 5 {
		t.Error("mergesort sort failed. Got,", result, "expect 1,2,3,4,5")
	}
}
func TestInsertSort2(t *testing.T) {
	values := []int{5, 4, 3, 2, 1, 11, 25, 44, 33, 16, 77}
	result := MergeSort(values)
	re := []int{1, 2, 3, 4, 5, 11, 16, 25, 33, 44, 77}
	for i := 0; i < len(re); i++ {
		if result[i] != re[i] {
			t.Error("mergesort sort failed. Got,", result, "expect", re)
		}
	}
}
