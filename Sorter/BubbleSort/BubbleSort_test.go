//bubblesort test
package bubblesort

import "testing"

func TestBubbleSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("bubblesort sort failed. Got,", values, "expect 1,2,3,4,5")
	}
}

func TestBubbleSortNew1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSortNew(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("bubblesort sort failed. Got,", values, "expect 1,2,3,4,5")
	}
}

/*
func TestSelection1New2(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	results := SelectionNew1(values)
	if results[0] != 1 || results[1] != 2 || results[2] != 3 || results[3] != 4 || results[4] != 5 {
		t.Error("bubblesort sort failed. Got,", results, "expect 1,2,3,4,5")
	}
}
*/
