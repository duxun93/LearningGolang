//ShellSort test
package shellsort

import "testing"

func TestShellSort1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	ShellSort(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("shellsort sort failed. Got,", values, "expect 1,2,3,4,5")
	}
}
