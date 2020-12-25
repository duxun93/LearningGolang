//test stack by queue
package stackAndqueue

import (
	"fmt"
	"testing"
)

func TestStackByQueue(t *testing.T) {
	expect := []int{-1, 3, 4, 2, 1}
	result := []int{}
	stack := NewStackByQueue()
	result = append(result, stack.Pop())
	fmt.Printf("pop")
	stack.Push(1)
	fmt.Printf("push")
	stack.Push(2)
	fmt.Printf("push")
	stack.Push(3)
	fmt.Printf("push")
	result = append(result, stack.Pop())
	fmt.Printf("pop")
	stack.Push(4)
	fmt.Printf("push")
	result = append(result, stack.Pop())
	fmt.Printf("pop")
	result = append(result, stack.Pop())
	fmt.Printf("pop")
	result = append(result, stack.Pop())
	fmt.Printf("pop")
	for index, value := range result {
		if value != expect[index] {
			t.Error("err should be", expect[index], "is", value)
		}
	}
}
