//test queue by stack
package stackAndqueue

import (
	"fmt"
	"testing"
)

func TestQueueByStack(t *testing.T) {
	expect := []int{-1, 1, 2, 3, 4}
	result := []int{}
	queue := NewQueueByStack()
	result = append(result, queue.Pop())
	fmt.Printf("pop")
	queue.Push(1)
	fmt.Printf("push")
	queue.Push(2)
	fmt.Printf("push")
	queue.Push(3)
	fmt.Printf("push")
	result = append(result, queue.Pop())
	fmt.Printf("pop")
	queue.Push(4)
	fmt.Printf("push")
	result = append(result, queue.Pop())
	fmt.Printf("pop")
	result = append(result, queue.Pop())
	fmt.Printf("pop")
	result = append(result, queue.Pop())
	fmt.Printf("pop")
	for index, value := range result {
		if value != expect[index] {
			t.Error("err should be", expect[index], "is", value)
		}
	}
}

/*
func TestBubbleSortNew1(t *testing.T) {
	values := []int{5, 4, 3, 2, 1}
	BubbleSortNew(values)
	if values[0] != 1 || values[1] != 2 || values[2] != 3 || values[3] != 4 || values[4] != 5 {
		t.Error("bubblesort sort failed. Got,", values, "expect 1,2,3,4,5")
	}
}
*/
