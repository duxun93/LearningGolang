//使用两个栈实现一个队列

package stackAndqueue

import (
	"fmt"
)

type QueueByStack struct {
	stack1 Stack
	stack2 Stack
}

func (rer *QueueByStack) Pop() (num int) {
	if len(rer.stack1.A) == 0 && len(rer.stack2.A) == 0 {
		return -1
	}
	if len(rer.stack2.A) == 0 {
		for len(rer.stack1.A) > 0 {
			temp := rer.stack1.Pop()
			rer.stack2.Push(temp)
		}
		fmt.Printf("stack2 ready")
	}
	return rer.stack2.Pop()
}

func (rer *QueueByStack) Push(num int) {
	rer.stack1.Push(num)
	return
}
func NewQueueByStack() *QueueByStack {
	result := new(QueueByStack)
	result.stack1 = *NewStack()
	result.stack2 = *NewStack()
	return result
}
